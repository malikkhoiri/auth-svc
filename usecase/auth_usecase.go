package usecase

import (
	"context"
	"time"

	"github.com/malikkhoiri/auth-svc/domain"
	"github.com/malikkhoiri/auth-svc/helper"
)

type authUsecase struct {
	authRepo   domain.AuthRepository
	ctxTimeout time.Duration
}

func NewAuthUsecase(authRepo domain.AuthRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		authRepo:   authRepo,
		ctxTimeout: timeout,
	}
}

func (au *authUsecase) LoginByEmailAndPassword(c context.Context, email, password string) (helper.M, error) {
	ctx, cancel := context.WithTimeout(c, au.ctxTimeout)
	defer cancel()

	res, err := au.authRepo.LoginByEmailAndPassword(ctx, email, password)

	return helper.M{"token": res}, err
}
