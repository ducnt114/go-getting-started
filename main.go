package main

import (
	"context"
	"fmt"
	"go-getting-started/cmd"
	"go-getting-started/pkg/log"
	"os"
)

func init() {
}

func main() {
	pid := os.Getpid()
	log.Infow(context.Background(), fmt.Sprintf("Process ID: %v", pid))
	cmd.Execute()
}
