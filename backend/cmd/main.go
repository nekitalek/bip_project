package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nekitalek/bip_project/backend/internal/repository"
	"github.com/nekitalek/bip_project/backend/internal/service"
	"github.com/nekitalek/bip_project/backend/internal/transport/handler"
	"github.com/nekitalek/bip_project/backend/internal/transport/http_server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title BIP API
// @version 1.0
// @description This is a sample serice for managing orders
// @contact.email kolya@example.com
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	fmt.Println("start app")
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	postres, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	db := repository.NewRepository(postres)
	services := service.NewService(db)
	handlers := handler.NewHandler(services)

	server := new(http_server.Server)
	if err := server.Run(viper.GetString("https.port"), viper.GetString("https.cert"), viper.GetString("https.key"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error starting server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
