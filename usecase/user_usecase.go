package usecase

import (
	"context"
	"time"

	"github.com/malikkhoiri/auth-svc/domain"
)

type userUsecase struct {
	userRepo   domain.UserRepository
	ctxTimeout time.Duration
}

func NewUserUsecase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		ctxTimeout: timeout,
	}
}

func (uc *userUsecase) Fetch(c context.Context, cursor string) ([]domain.User, string, error) {
	users := make([]domain.User, 0)
	return users, "", nil
}

func (uc *userUsecase) GetByID(c context.Context, id int64) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *userUsecase) Store(c context.Context, u *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	return uc.userRepo.Store(ctx, u)
}

func (uc *userUsecase) Update(c context.Context, u *domain.User) error {
	return nil
}
