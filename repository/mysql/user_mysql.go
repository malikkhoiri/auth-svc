package mysql

import (
	"context"
	"database/sql"

	"github.com/malikkhoiri/auth-svc/domain"
	"golang.org/x/crypto/bcrypt"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func NewMysqlUserRepository(Conn *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn: Conn}
}

func (m *mysqlUserRepository) Fetch(ctx context.Context, cursor string) ([]domain.User, string, error) {
	query := "SELECT id, name, username, email, created_at, updated_at FROM users ORDER BY created_at DESC"
	rows, err := m.Conn.QueryContext(ctx, query)

	if err != nil {
		return nil, "", err
	}

	users := make([]domain.User, 0)
	for rows.Next() {
		user := domain.User{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, "", err
		}

		users = append(users, user)
	}

	return users, "", nil
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	query := "SELECT id, name, username, email, created_at, updated_at FROM users WHERE id=?"
	stmt, err := m.Conn.Prepare(query)
	user := domain.User{}

	if err != nil {
		return user, err
	}

	err = stmt.QueryRow(id).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
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

	stmt.ExecContext(ctx, u.Name, u.Username, u.Email, string(crypt))

	return nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, u *domain.User) error {
	return nil
}
