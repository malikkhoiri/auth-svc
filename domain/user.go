package domain

import (
	"context"
	"time"

	"github.com/malikkhoiri/auth-svc/helper"
)

type User struct {
	ID        int               `json:"id"`
	Name      string            `json:"name" validate:"required"`
	Username  helper.NullString `json:"username"`
	Email     string            `json:"email" validate:"required"`
	Password  string            `json:"password,omitempty" validate:"required"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type UserRepository interface {
	Fetch(ctx context.Context, cursor string) ([]User, string, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Store(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
}

type UserUsecase interface {
	Fetch(ctx context.Context, cursor string) ([]User, string, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Store(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
}
