package main

import (
	"go-getting-started/cmd"
	"go.uber.org/zap"
	"os"
)

func init() {
}

func main() {
	pid := os.Getpid()
	zap.S().Infof("Process ID: %v", pid)
	cmd.Execute()
}
