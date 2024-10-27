package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
	"go-getting-started/dto"
	"go-getting-started/log"
	"go-getting-started/repository"
	"go-getting-started/utils"
	"net/http"
	"time"
)

type AuthService interface {
	PasswordLogin(ctx context.Context, req *dto.PasswordLoginRequest) (*dto.LoginResponse, error)
}

type authServiceImpl struct {
	userRepo repository.UserRepository
	jwtUtil  utils.JWTUtil
}

func newAuthService(di *do.Injector) (AuthService, error) {
	return &authServiceImpl{
		userRepo: do.MustInvoke[repository.UserRepository](di),
		jwtUtil:  do.MustInvoke[utils.JWTUtil](di),
	}, nil
}

func (s *authServiceImpl) PasswordLogin(
	ctx context.Context,
	req *dto.PasswordLoginRequest,
) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	// TODO validate hashed password
	if user.Pass != req.Password {
		return nil, errors.New("invalid password")
	}

	// generate access token
	currentTime := time.Now()
	claims := dto.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
		UserName: user.Name,
		UserUUID: fmt.Sprintf("%v", user.ID), // TODO use uuid
	}

	accessToken, err := s.jwtUtil.GenerateToken(&claims)
	if err != nil {
		log.Errorw(ctx, "error when generating token for user", "err", err)
		return nil, err
	}

	return &dto.LoginResponse{
		Meta: &dto.Meta{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: &dto.Token{
			AccessToken: accessToken,
		},
	}, nil
}
