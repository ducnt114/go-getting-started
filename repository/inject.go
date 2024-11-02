package repository

import "github.com/samber/do"

func Inject(di *do.Injector) {
	do.Provide(di, newUserRepository)
	do.Provide(di, newTokenRepository)
	do.Provide(di, newTwoFaRepository)
}
