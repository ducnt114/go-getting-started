package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"go-getting-started/dto"
	"go-getting-started/log"
	"go-getting-started/model"
	"go-getting-started/repository"
	"gorm.io/gorm"
	"net/http"
)

type UserService interface {
	GetUserById(ctx context.Context, userId uint) (*dto.User, error)
	CreateUser(ctx context.Context, req *dto.User) (*dto.User, error)
	PasswordLogin(ctx context.Context, username, password string) (*dto.PasswordLoginResponse, error)
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
	span := sentry.StartSpan(ctx, "userServiceImpl.GetUserById")
	span.Description = "GetUserById_service"
	defer span.Finish()

	user, err := u.userRepo.FindByID(span.Context(), userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with id: %d", userId)
		}
		return nil, err
	}
	log.Infow(ctx, "get user by id", "user", user.Name)
	//time.Sleep(2 * time.Second)
	res := &dto.User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	return res, nil
}

func (u *userServiceImpl) CreateUser(ctx context.Context, req *dto.User) (*dto.User, error) {
	existedUser, _ := u.userRepo.FindByName(ctx, req.Name)
	if existedUser != nil && existedUser.Name != "" {
		return nil, errors.New("user existed")
	}
	user := &model.User{
		Name: req.Name,
		Age:  req.Age,
	}
	// validate password
	// ....
	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (u *userServiceImpl) PasswordLogin(
	ctx context.Context,
	username, password string,
) (*dto.PasswordLoginResponse, error) {

	user, err := u.userRepo.FindByName(ctx, username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if user.Pass != password {
		return nil, errors.New("invalid user/pass")
	}
	return &dto.PasswordLoginResponse{
		Meta: &dto.Meta{
			Code:    http.StatusOK,
			Message: user.Name,
		},
	}, nil
}
