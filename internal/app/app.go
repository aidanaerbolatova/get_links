package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"test/config"
	"test/internal/app/logger"
	"test/internal/repository"
	"test/internal/service"
	test "test/internal/transport/httpserver"
	"test/internal/transport/httpserver/handlers"
)

func Run() error {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)
	logger, err := logger.Logger()
	if err != nil {
		logger.Errorf("Error while initialization logger: %v", err)
		return err
	}
	config, err := config.ReadConfig()
	if err != nil {
		logger.Errorf("Error while parse config file: %v", err)
		return err
	}
	fmt.Println(config)
	db, err := repository.ConnectDB(logger, config)
	if err != nil {
		logger.Errorf("Error while connect to DB: %v", err)
		return err
	}
	redis, err := repository.ConnectRedis(logger, config)
	if err != nil {
		logger.Errorf("Error while connect to redis: %v", err)
		return err
	}
	repository := repository.NewRepository(db, redis, logger)
	service := service.NewService(repository, logger)
	handlers := handlers.NewHandler(service)

	_, err = service.AddToDB()
	if err != nil {
		logger.Errorf("Error while add dates to DB: %v", err)
		return err
	}

	// handlers.CacheWarming(data)

	srv := new(test.Server)
	if err := srv.Run("8080", handlers.InitRoute()); err != nil {
		logger.Errorf("Error while start server: %v", err)
		return err
	}

	return nil
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
