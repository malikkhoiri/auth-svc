package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/malikkhoiri/auth-svc/helper"
)

type Auth struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"Username"`
}

type AuthRepository interface {
	LoginByEmailAndPassword(ctx context.Context, username, email string) (token string, err error)
}

type AuthUsecase interface {
	LoginByEmailAndPassword(ctx context.Context, email, password string) (helper.M, error)
}
