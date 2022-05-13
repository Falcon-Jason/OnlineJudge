package service

import (
	"OnlineJudge_JudgerServer/pb/judge_result"
	"OnlineJudge_JudgerServer/pb/submission"
	"context"
)

type judgeService struct {
	submission.UnimplementedJudgeServer
	wbClient judge_result.ResultWriteBackClient
}

func NewJudgeService(client judge_result.ResultWriteBackClient) *judgeService {
	return &judgeService{wbClient: client}
}

func (s *judgeService) Judge(ctx context.Context, req *submission.Submission) (*submission.JudgeReply, error) {
	go s.judge(req)
	return &submission.JudgeReply{Ok: true}, nil
}
