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

	httpDelivery "github.com/malikkhoiri/auth-svc/delivery/http"
	"github.com/malikkhoiri/auth-svc/helper"
	repo "github.com/malikkhoiri/auth-svc/repository/mysql"
	usecase "github.com/malikkhoiri/auth-svc/usecase"
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
	// database config
	dbName := viper.GetString(`mysql.database`)
	mysqlHost := viper.GetString(`mysql.host`)
	mysqlPort := viper.GetString(`mysql.port`)
	mysqlUser := viper.GetString(`mysql.username`)
	mysqlPass := viper.GetString(`mysql.password`)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysqlUser, mysqlPass, mysqlHost, mysqlPort, dbName)
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
	timeoutCtx := time.Duration(2) * time.Second

	// Authentication
	authRepo := repo.NewMysqlAuthRepository(dbConn)
	au := usecase.NewAuthUsecase(authRepo, timeoutCtx)
	httpDelivery.NewAuthHandler(e, au)

	// User
	userRepo := repo.NewMysqlUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepo, timeoutCtx)
	httpDelivery.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start(viper.GetString("server.port")))
}
