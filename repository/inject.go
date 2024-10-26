package repository

import "github.com/samber/do"

func Inject(di *do.Injector) {
	do.Provide(di, NewUserRepository)
}
