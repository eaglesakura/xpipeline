package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"strings"
	"runtime"
)

/*
 Dockerで必要なPathに変換するタスク
*/
type DockerPathTask struct {
	ctx *cli.Context
}

func newDockerPathTask(ctx *cli.Context) (*DockerPathTask, error) {
	return &DockerPathTask{
		ctx: ctx,
	}, nil
}

func (it *DockerPathTask) Execute() {
	currentDir, _ := os.Getwd()

	if runtime.GOOS == "windows" {
		if currentDir[0:1] != "/" {
			currentDir = "/" + strings.ToLower(currentDir[0:1]) + currentDir[1:]
		}
		currentDir = strings.Replace(currentDir, ":\\", "/", -1)
		currentDir = strings.Replace(currentDir, "\\", "/", -1)
	}
	fmt.Print(currentDir)
}
