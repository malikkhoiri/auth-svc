package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/malikkhoiri/auth-svc/helper"
)

type Auth struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"Username"`
}

type AuthRepository interface {
	LoginByUsernameAndPassword(ctx context.Context, username, password string) (token string, err error)
}

type AuthUsecase interface {
	LoginByUsernameAndPassword(ctx context.Context, username, password string) (helper.M, error)
}
