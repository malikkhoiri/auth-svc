package mysql

import (
	"context"
	"database/sql"

	"github.com/malikkhoiri/auth-svc/domain"
	"github.com/malikkhoiri/auth-svc/helper"
	"golang.org/x/crypto/bcrypt"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func NewMysqlUserRepository(Conn *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn: Conn}
}

func (m *mysqlUserRepository) Fetch(ctx context.Context, cursor string) ([]domain.User, string, error) {
	users := make([]domain.User, 0)
	return users, "", nil
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return domain.User{}, nil
}

func (m *mysqlUserRepository) Store(ctx context.Context, u *domain.User) error {
	query := "INSERT users (name, username, email, password) VALUES (?, ?, ?, ?)"
	stmt, err := m.Conn.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	crypt, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	stmt.ExecContext(ctx, u.Name, helper.NullString(u.Username), u.Email, string(crypt))

	return nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, u *domain.User) error {
	return nil
}
