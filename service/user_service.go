package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/samber/do"
	"go-getting-started/dto"
	"go-getting-started/model"
	"go-getting-started/repository"
	"go-getting-started/utils"
	"gorm.io/gorm"
	"net/http"
)

const (
	saltLength = 20
)

type UserService interface {
	GetUserById(ctx context.Context, userId uint) (*dto.User, error)
	List(ctx context.Context, name string) (*dto.ListUserResponse, error)
	Update(ctx context.Context, id uint) (*dto.User, error)
	CreateUser(ctx context.Context, req *dto.CreateUserReq) (*dto.User, error)
	PasswordLogin(ctx context.Context, username, password string) (*dto.PasswordLoginResponse, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func newUserService(di *do.Injector) (UserService, error) {
	return &userServiceImpl{
		userRepo: do.MustInvoke[repository.UserRepository](di),
	}, nil
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
		Pass: req.Password,
	}

	salt := utils.RandomStringWithLength(saltLength)
	hashedPass := utils.HashPassword(req.Password, salt)
	user.Pass = hashedPass
	user.Salt = salt

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	userRes := s.convertToUserDto(user)
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
	if user.Tags != nil && len(user.Tags) > 0 {
		res.Tag1 = user.Tags[0].Val
	}
	return res
}

func (s *userServiceImpl) Update(ctx context.Context, id uint) (*dto.User, error) {
	u, err := s.userRepo.UpdateUserAgeDemo(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.convertToUserDto(u), nil
}
