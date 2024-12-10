package main

import (
	_ "github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/mickey-mickser/simplecoin"
	"github.com/mickey-mickser/simplecoin/pkg/handler"
	"github.com/mickey-mickser/simplecoin/pkg/repository"
	"github.com/mickey-mickser/simplecoin/pkg/usecase"
	"github.com/mickey-mickser/simplecoin/runner"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

func main() {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("failed to init config: %v", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error write .env: %s", err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to init db: %v", err.Error())
	}

	//зависимсти
	repos := repository.NewRepository(db)
	services := usecase.NewPricesUseCase(repos)
	handlers := handler.NewHandler(services)

	srv := new(crypto.Server)
	go func() {
		if err := srv.Start(viper.GetString("port"), handlers.InitRouter()); err != nil {
			logrus.Fatalf("error occurred while running the HTTP server: %s", err.Error())
		}
	}()

	rep := repository.NewPricesPostgres(db)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := runner.ParseCoins(rep); err != nil {
			logrus.Fatalf("failed to parse coins: %v", err.Error())
		}
	}()
	wg.Wait()

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
