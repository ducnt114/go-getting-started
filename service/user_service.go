package service

import (
	"context"
	"go-getting-started/dto"
	"go-getting-started/log"
	"go-getting-started/model"
	"go-getting-started/repository"
)

type UserService interface {
	GetUserById(ctx context.Context, userId uint) (*dto.User, error)
	CreateUser(ctx context.Context, req *dto.User) (*dto.User, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (u *userServiceImpl) GetUserById(ctx context.Context, userId uint) (*dto.User, error) {
	user, err := u.userRepo.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	log.Infow(ctx, "get user by id", "user", user.Name)
	res := &dto.User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	return res, nil
}

func (u *userServiceImpl) CreateUser(ctx context.Context, req *dto.User) (*dto.User, error) {
	user := &model.User{
		Name: req.Name,
		Age:  req.Age,
	}
	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return req, nil
}
