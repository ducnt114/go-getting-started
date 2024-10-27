package utils

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
	"go-getting-started/conf"
	"os"
)

// RS256 is RSA Signature with SHA-256.
const RS256 = "RS256"

// JWTUtil decodes or encodes JWT access token.
type JWTUtil interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ParseClaims(token string, claims jwt.Claims) error
}

type jwtUtil struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewJWTUtil returns a new instance of JWTHelper.
func NewJWTUtil(di *do.Injector) (JWTUtil, error) {
	cf := do.MustInvoke[*conf.Config](di)
	jwtPubKey, err := os.ReadFile(cf.JWT.PublicKeyFilePath)
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtPubKey)
	if err != nil {
		panic(err)
	}

	jwtPrivateKey, err := os.ReadFile(cf.JWT.PrivateKeyFilePath)
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	jwtUtl := &jwtUtil{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	return jwtUtl, nil
}

func (h *jwtUtil) GenerateToken(claims jwt.Claims) (string, error) {
	if claims == nil {
		return "", errors.New("claims must not be nil")
	}

	tkn := jwt.NewWithClaims(jwt.GetSigningMethod(RS256), claims)
	str, err := tkn.SignedString(h.privateKey)
	if err != nil {
		return "", err
	}

	return str, nil
}

func (h *jwtUtil) ParseClaims(tokenStr string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return h.publicKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
