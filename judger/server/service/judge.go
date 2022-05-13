package service

import (
	"OnlineJudge_JudgerServer/judger"
	"OnlineJudge_JudgerServer/pb/judge_result"
	pb "OnlineJudge_JudgerServer/pb/judger"
	"OnlineJudge_JudgerServer/pb/submission"
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	judgerUid   = flag.Int("judger_uid", 1001, "The user id for judger")
	workspace   = flag.String("workspace", "/home/judge/problem", "the path of workspace")
	environment = os.Environ()
)

const (
	MAX_STACK       = 128 * 1024 * 1024
	MAX_OUTPUT_SIZE = 20 * 1024 * 1024
)

func (s *judgeService) writeBack(req *submission.Submission, status judge_result.JudgeStatus) {
	_, _ = s.wbClient.WriteBack(context.TODO(), &judge_result.ResultWriteBackRequest{
		Id:     req.Id,
		Status: status,
	})
}

func (s *judgeService) judge(req *submission.Submission) {
	ws := path.Join(*workspace, req.Id)
	err := os.MkdirAll(ws, os.FileMode(0755))

	if err != nil {
		log.Print(err)
		s.writeBack(req, judge_result.JudgeStatus_internal_error)
		return
	}
	err = os.Chown(ws, *judgerUid, *judgerUid)
	if err != nil {
		log.Print(err)
		s.writeBack(req, judge_result.JudgeStatus_internal_error)
		return
	}

	lcg, ok := languageConfigs[req.CodeLanguage]
	if !ok {
		s.writeBack(req, judge_result.JudgeStatus_internal_error)
		return
	}

	lc := lcg(ws)

	err = ioutil.WriteFile(lc.cc.srcName, []byte(req.CodeText), os.FileMode(0644))
	if err != nil {
		log.Print(err)
		s.writeBack(req, judge_result.JudgeStatus_internal_error)
		return
	}

	err = os.Chown(lc.cc.srcName, *judgerUid, *judgerUid)
	if err != nil {
		log.Print(err)
		s.writeBack(req, judge_result.JudgeStatus_internal_error)
		return
	}

	s.writeBack(req, judge_result.JudgeStatus_compiling)
	ok = s.compile(ws, &lc.cc)
	if !ok {
		_, _ = s.wbClient.WriteBack(context.TODO(), &judge_result.ResultWriteBackRequest{
			Id:     req.Id,
			Status: judge_result.JudgeStatus_compile_error,
		})
		return
	}

	s.writeBack(req, judge_result.JudgeStatus_running)
	for _, tc := range req.TestCases {
		status := s.runAndCompare(ws, &lc.rc, tc)
		if status != judge_result.JudgeStatus_accept {
			s.writeBack(req, status)
			return
		}
	}

	s.writeBack(req, judge_result.JudgeStatus_accept)
}

func (s *judgeService) compile(ws string, c *compileConfig) bool {
	result, err := judger.Run(&pb.Config{
		MaxCpuTime:       c.maxCpuTime,
		MaxRealTime:      c.maxRealTime,
		MaxMemory:        c.maxMemory,
		MaxStack:         MAX_STACK,
		MaxOutputSize:    MAX_OUTPUT_SIZE,
		MaxProcessNumber: -1,
		ExePath:          c.command[0],
		InputPath:        "/dev/null",
		OutputPath:       path.Join(ws, "compile_out"),
		ErrorPath:        path.Join(ws, "compile_out"),
		Args:             c.command[1:],
		Env:              environment,
		LogPath:          path.Join(*workspace, "log/compile.log"),
		Uid:              uint32(*judgerUid),
		Gid:              int32(*judgerUid),
		UseSeccomp:       false,
	})

	log.Print("compile success")
	if err != nil {
		log.Print(err)
		return false
	}

	if result.Error != pb.ErrorCode_SUCCESS {
		log.Print(result.Error.String(), result.Signal, result.Result)
		return false
	}

	if _, err := os.Stat(c.exeName); os.IsNotExist(err) {
		log.Print(err)
		log.Print(c.command)
		return false
	}

	return true
}

func (s *judgeService) runAndCompare(ws string, c *runConfig, tc *submission.TestCase) judge_result.JudgeStatus {
	err := ioutil.WriteFile(path.Join(ws, "in"), []byte(tc.Input), 0)
	if err != nil {
		return judge_result.JudgeStatus_internal_error
	}

	result, err := judger.Run(&pb.Config{
		MaxCpuTime:       c.maxCpuTime,
		MaxRealTime:      c.maxRealTime,
		MaxMemory:        c.maxMemory,
		MaxStack:         MAX_STACK,
		MaxOutputSize:    MAX_OUTPUT_SIZE,
		MaxProcessNumber: -1,
		ExePath:          c.command[0],
		InputPath:        path.Join(ws, "in"),
		OutputPath:       path.Join(ws, "out"),
		ErrorPath:        "/dev/null",
		Args:             c.command[1:],
		Env:              []string{},
		LogPath:          path.Join(*workspace, "log/run.log"),
		Uid:              uint32(*judgerUid),
		Gid:              int32(*judgerUid),
		UseSeccomp:       true,
	})

	if err != nil {
		log.Print(err)
		return judge_result.JudgeStatus_internal_error
	}

	switch result.Result {
	case pb.ResultCode_CPU_TIME_LIMIT_EXCEEDED:
	case pb.ResultCode_REAL_TIME_LIMIT_EXCEEDED:
		return judge_result.JudgeStatus_time_limited_error
	case pb.ResultCode_MEMORY_LIMIT_EXCEEDED:
		return judge_result.JudgeStatus_memory_limited_error
	case pb.ResultCode_RUNTIME_ERROR:
		log.Print(result.Signal, result.ExitCode)
		return judge_result.JudgeStatus_runtime_error
	case pb.ResultCode_SYSTEM_ERROR:
		return judge_result.JudgeStatus_internal_error
	case pb.ResultCode_WRONG_ANSWER:
		return judge_result.JudgeStatus_wrong_answer
	case pb.ResultCode_ACCEPTED:
		// no return
	}

	if result.CpuTime > c.maxCpuTime || result.RealTime > c.maxRealTime {
		return judge_result.JudgeStatus_time_limited_error
	}

	data, err := ioutil.ReadFile(path.Join(ws, "out"))
	if err != nil {
		return judge_result.JudgeStatus_internal_error
	}

	if string(data) == tc.Output {
		return judge_result.JudgeStatus_accept
	} else {
		return judge_result.JudgeStatus_wrong_answer
	}
}
