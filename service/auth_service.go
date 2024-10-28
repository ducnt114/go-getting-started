package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
	"go-getting-started/dto"
	"go-getting-started/log"
	"go-getting-started/model"
	"go-getting-started/repository"
	"go-getting-started/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type AuthService interface {
	PasswordLogin(ctx context.Context, req *dto.PasswordLoginRequest) (*dto.LoginResponse, error)
	GetTwoFa(ctx context.Context, userID uint) (*dto.GetTwoFaResponse, error)
	SetupTwoFa(ctx context.Context, req *dto.SetupTwoFaRequest) (*dto.SetupTwoFaResponse, error)
}

type authServiceImpl struct {
	userRepo  repository.UserRepository
	twoFaRepo repository.TwoFaRepository
	jwtUtil   utils.JWTUtil
}

func newAuthService(di *do.Injector) (AuthService, error) {
	return &authServiceImpl{
		userRepo:  do.MustInvoke[repository.UserRepository](di),
		twoFaRepo: do.MustInvoke[repository.TwoFaRepository](di),
		jwtUtil:   do.MustInvoke[utils.JWTUtil](di),
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
		UserID:   user.ID,
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

func (s *authServiceImpl) GetTwoFa(ctx context.Context, userID uint) (*dto.GetTwoFaResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user.TwoFA != "" {
		return nil, errors.New("2fa already setup")
	}
	oldTwoFa, err := s.twoFaRepo.FindByUserID(ctx, user.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newSecret := utils.GenerateTwoFaSecret()
		err = s.twoFaRepo.Save(ctx, &model.TwoFactor{
			UserID: user.ID,
			Secret: newSecret})
		if err != nil {
			log.Errorw(ctx, "Error when save temp 2fa secret", "err", err)
			return &dto.GetTwoFaResponse{Meta: dto.InternalServerErrorMeta}, nil
		}
		return &dto.GetTwoFaResponse{
			Meta: dto.SuccessMeta,
			Data: &dto.TwoFaResponseData{
				Secret: newSecret,
			},
		}, nil
	}
	// return generated 2fa secret
	return &dto.GetTwoFaResponse{
		Meta: dto.SuccessMeta,
		Data: &dto.TwoFaResponseData{
			Secret: oldTwoFa.Secret,
		},
	}, nil
}

func (s *authServiceImpl) SetupTwoFa(
	ctx context.Context,
	req *dto.SetupTwoFaRequest,
) (*dto.SetupTwoFaResponse, error) {
	return nil, nil // TODO
}
