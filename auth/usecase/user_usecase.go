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

func (uc *userUsecase) Fetch(ctx context.Context, cursor string) ([]domain.User, string, error) {
	users := make([]domain.User, 0)
	return users, "", nil
}

func (uc *userUsecase) GetByID(ctx context.Context, id int64) (domain.User, error) {
	user := domain.User{}
	return user, nil
}

func (uc *userUsecase) Store(ctx context.Context, u *domain.User) error {
	return nil
}

func (uc *userUsecase) Update(ctx context.Context, u *domain.User) error {
	return nil
}
