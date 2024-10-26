package api

import (
	"context"
	"fmt"
	"github.com/samber/do"
	"github.com/spf13/cobra"
	"go-getting-started/conf"
	"go-getting-started/connection"
	"go-getting-started/log"
	"go-getting-started/repository"
	"go-getting-started/service"
)

var Cmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  `api`,
	Run: func(cmd *cobra.Command, args []string) {
		startApi()
	},
}

func startApi() {
	injector := do.New()
	defer func() {
		_ = injector.Shutdown()
	}()
	conf.Inject(injector)
	connection.Inject(injector)
	repository.Inject(injector)
	service.Inject(injector)

	r, err := InitRouter(injector)
	if err != nil {
		panic(err)
	}

	cf := do.MustInvoke[*conf.Config](injector)
	addr := fmt.Sprintf(":%v", cf.ApiService.Port)
	log.Infow(context.Background(), fmt.Sprintf("start api server at %v", addr))
	_ = r.Run(addr)
}
