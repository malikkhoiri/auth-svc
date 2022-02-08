package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/malikkhoiri/auth-svc/domain"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var JWT_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type mysqlAuthRepository struct {
	Conn *sql.DB
}

func NewMysqlAuthRepository(Conn *sql.DB) domain.AuthRepository {
	return &mysqlAuthRepository{Conn}
}

func (m *mysqlAuthRepository) LoginByEmailAndPassword(ctx context.Context, email, password string) (result string, err error) {
	query := "SELECT email, password FROM user WHERE email = ?"
	stmt, err := m.Conn.Prepare(query)

	if err != nil {
		return
	}

	auth := domain.Auth{}
	err = stmt.QueryRow(email).Scan(
		&auth.Email,
		&auth.Password,
	)

	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))

	if err != nil {
		return
	}

	claims := domain.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("serviceName"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRATION_DURATION)),
		},
		Email: email,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	result, err = token.SignedString([]byte(viper.GetString(`jwt.signatureKey`)))

	if err != nil {
		return
	}

	return
}
