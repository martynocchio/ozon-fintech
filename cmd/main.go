package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"ozon-fintech/pkg/handler"
	"ozon-fintech/pkg/repository/inmemory"
	"ozon-fintech/pkg/repository/postgres"
	"ozon-fintech/pkg/service"
)

const (
	databaseURLKey = "DATABASE_URL"
	port           = "port"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %v", err)
	}

	var dbFlag bool
	flag.BoolVar(&dbFlag, "db", false, "Run with DB postgres:")
	flag.Parse()
	logrus.SetFormatter(new(logrus.JSONFormatter))

	var repos service.Services
	if dbFlag {
		dbConfig := os.Getenv(databaseURLKey)
		if dbConfig == "" {
			logrus.Info("empty env config")
			dbConfig = viper.GetString(databaseURLKey)
		}

		db, err := postgres.NewPostgresDB(postgres.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		})
		if err != nil {
			logrus.Fatalf("failed to initialize db: %v", err)
		}
		repos = postgres.NewRepository(db)
	} else {
		repos = inmemory.NewRepository()
	}

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	port := viper.GetString(port)

	app := echo.New()
	handlers.InitRotes(app)

	if err := app.Start(":" + port); err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
