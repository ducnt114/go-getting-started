package dto

import "github.com/golang-jwt/jwt/v5"

type LoginResponse struct {
	Meta *Meta  `json:"meta"`
	Data *Token `json:"data"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserUUID string `json:"user_uuid"`
	UserName string `json:"user_name"`
}

type PasswordLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
