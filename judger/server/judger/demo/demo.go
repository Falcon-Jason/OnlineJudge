package main

import (
	"OnlineJudge_JudgerServer/judger"
	pb "OnlineJudge_JudgerServer/pb/judger"
	"fmt"
	"io/ioutil"
	"os/exec"
)

const (
	inputFile = "World\n"

	sourceFile = `
#include <stdio.h>
int main(int argc, char *argv[]) {
    char input[1000];
    scanf("%s", input);
    printf("Hello %s\n", input);
    return 0;
}`
)

func main() {
	cmd := exec.Command("gcc", "temp/main.c", "-o", "temp/main")
	_ = cmd.Start()
	_ = cmd.Wait()

	_ = ioutil.WriteFile("temp/1.in", []byte(inputFile), 0)
	_ = ioutil.WriteFile("temp/main.c", []byte(sourceFile), 0)

	result, err := (&judger.Judger{"../../../client/build/judge"}).
		Run(&pb.Config{
			MaxCpuTime:       1000,
			MaxRealTime:      2000,
			MaxMemory:        128 * 1024 * 1024,
			MaxProcessNumber: 200,
			MaxOutputSize:    10000,
			MaxStack:         32 * 1024 * 1024,
			// five args above can be _judger.UNLIMITED
			ExePath:    "temp/main",
			InputPath:  "temp/1.in",
			OutputPath: "temp/1.out",
			ErrorPath:  "temp/1.out",
			Args:       []string{},
			// can be empty list
			Env:     []string{},
			LogPath: "judger.log",
			// can be None
			Uid: 0,
			Gid: 0,
		})

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.Error, result.Result)
}
