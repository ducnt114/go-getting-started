package service

import (
	"context"
	"errors"
	"fmt"
	"go-getting-started/dto"
	"go-getting-started/model"
	"go-getting-started/repository"
	"gorm.io/gorm"
	"net/http"
)

type UserService interface {
	GetUserById(ctx context.Context, userId uint) (*dto.User, error)
	List(ctx context.Context, name string) (*dto.ListUserResponse, error)
	CreateUser(ctx context.Context, req *dto.CreateUserReq) (*dto.User, error)
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

func (s *userServiceImpl) GetUserById(ctx context.Context, userId uint) (*dto.User, error) {
	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with id: %d", userId)
		}
		return nil, err
	}
	res := s.convertToUserDto(user)
	return res, nil
}

func (s *userServiceImpl) CreateUser(ctx context.Context, req *dto.CreateUserReq) (*dto.User, error) {
	existedUser, _ := s.userRepo.FindByName(ctx, req.Name)
	if existedUser != nil && existedUser.Name != "" {
		return nil, errors.New("user existed")
	}
	user := &model.User{
		Name: req.Name,
		Age:  req.Age,
	}
	// validate password
	// ....
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	userRes := &dto.User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
		Bio:  "",
	}
	return userRes, nil
}

func (s *userServiceImpl) PasswordLogin(
	ctx context.Context,
	username, password string,
) (*dto.PasswordLoginResponse, error) {

	user, err := s.userRepo.FindByName(ctx, username)
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

func (s *userServiceImpl) List(ctx context.Context, name string) (*dto.ListUserResponse, error) {
	users, err := s.userRepo.List(ctx, name)
	if err != nil {
		return nil, err
	}
	listUser := make([]*dto.User, 0)
	for _, u := range users {
		listUser = append(listUser, s.convertToUserDto(u))
	}
	res := &dto.ListUserResponse{
		Meta: &dto.Meta{
			Code:    http.StatusOK,
			Message: "ok",
		},
		Data: listUser,
	}
	return res, nil
}

func (s *userServiceImpl) convertToUserDto(user *model.User) *dto.User {
	res := &dto.User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	if user.Profile != nil {
		res.Bio = user.Profile.Bio
	}
	for _, b := range user.Books {
		res.Books = append(res.Books, &dto.Book{
			Name:  b.Name,
			Title: b.Title,
		})
	}
	return res
}
