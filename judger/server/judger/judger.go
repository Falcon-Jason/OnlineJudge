package judger

import (
	pb "OnlineJudge_JudgerServer/pb/judger"
	"flag"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"os/exec"
)

var (
	judgerPath = flag.String("judger_path", "/usr/bin/judger", "The path of judger")
)

func Run(config *pb.Config) (*pb.Result, error) {
	cmd := exec.Command(*judgerPath)

	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()

	err := cmd.Start()
	log.Print("start cmd")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	data, err := proto.Marshal(config)
	log.Println("marshal config")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	_, err = in.Write(data)
	log.Println("write config")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	_ = in.Close()

	data, err = ioutil.ReadAll(out)
	log.Println("read output")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var result pb.Result
	err = proto.Unmarshal(data, &result)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &result, nil
}
