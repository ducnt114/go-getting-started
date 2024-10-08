package main

import (
	"go-getting-started/cmd"
	"go-getting-started/conf"
	"go.uber.org/zap"
	"os"
)

func init() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	pid := os.Getpid()
	zap.S().Infof("Process ID: %v", pid)
	cmd.Execute()
}
