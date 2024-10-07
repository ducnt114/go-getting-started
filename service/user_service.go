package service

import (
	"context"
	"go-getting-started/dto"
	"go-getting-started/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (u *UserService) GetUserById(ctx context.Context, userId uint) (*dto.User, error) {
	user, err := u.UserRepo.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	res := &dto.User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	return res, nil
}
