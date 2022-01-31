package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/malikkhoiri/auth-svc/domain"
	"github.com/spf13/viper"
)

var JWT_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type mysqlAuthRepository struct {
	Conn *sql.DB
}

func NewMysqlAuthRepository(Conn *sql.DB) domain.AuthRepository {
	return &mysqlAuthRepository{Conn}
}

func (m *mysqlAuthRepository) LoginByUsernameAndPassword(ctx context.Context, username, password string) (result string, err error) {

	claims := domain.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("serviceName"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRATION_DURATION)),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString(`jwt.signatureKey`)))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
