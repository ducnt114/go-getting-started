package utils

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/samber/do"
	"go-getting-started/conf"
	"go-getting-started/log"
	"gorm.io/gorm"
)

type Enforcer struct {
	E *casbin.Enforcer
}

func NewEnforcer(di *do.Injector) (*Enforcer, error) {
	cf := do.MustInvoke[*conf.Config](di)
	db := do.MustInvoke[*gorm.DB](di)
	adapter, err := gormAdapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	//e, err := casbin.NewEnforcer(cf.Casbin.RBACModelPath, cf.Casbin.RBACPolicyPath)
	e, err := casbin.NewEnforcer(cf.Casbin.RBACModelPath, adapter)
	if err != nil {
		log.Errorw(context.Background(), "error when init casbin", "err", err)
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("error in policy: %w", err)
	}

	return &Enforcer{
		E: e,
	}, nil
}
