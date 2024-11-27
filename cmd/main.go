package main

import (
	"os"
	restapi "rest_api_learn"
	"rest_api_learn/pgk/handler"
	"rest_api_learn/pgk/repository"
	"rest_api_learn/pgk/service"
	"rest_api_learn/utils"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	if err := utils.InitConfig(); err != nil {
		logrus.Fatalf("error initialization YAML config: %s \n", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initialization ENV config: %s \n", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:      viper.GetString("db.port"),
		Username:  viper.GetString("db.username"),
		Password:  os.Getenv("PG_DB_PASS"),
		DBName:    viper.GetString("db.dbname"),
		SSLMode:   viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error initioaization POSTGRES DATABASE: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	srv := new(restapi.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while run http server: %s", err.Error())
	}
}
