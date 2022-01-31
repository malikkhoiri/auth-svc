package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_authHttp "github.com/malikkhoiri/auth-svc/auth/delivery/http"
	_authRepo "github.com/malikkhoiri/auth-svc/auth/repository"
	_authUsecase "github.com/malikkhoiri/auth-svc/auth/usecase"
	"github.com/malikkhoiri/auth-svc/helper"
)

func init() {
	log.Println(time.Hour)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	dbName := viper.GetString(`mysql.database`)
	mysqlHost := viper.GetString(`mysql.host`)
	mysqlPort := viper.GetString(`mysql.port`)
	mysqlUser := viper.GetString(`mysql.username`)
	mysqlPass := viper.GetString(`mysql.password`)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, dbName)
	dbConn, err := sql.Open("mysql", conn)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	authRepo := _authRepo.NewMysqlAuthRepository(dbConn)
	timeoutCtx := time.Duration(2) * time.Second
	au := _authUsecase.NewAuthUsecase(authRepo, timeoutCtx)
	_authHttp.NewAuthHandler(e, au)
	log.Fatal(e.Start(viper.GetString("server.port")))
}
