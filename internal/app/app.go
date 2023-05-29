package app

import (
	"test/config"
	"test/internal/app/logger"
	"test/internal/repository"
	"test/internal/service"
	test "test/internal/transport/httpserver"
	"test/internal/transport/httpserver/handlers"
)

func Run() error {
	logger, err := logger.Logger()
	if err != nil {
		logger.Errorf("Error while initialization logger: %v", err)
		return err
	}
	config, err := config.ParseYaml()
	if err != nil {
		logger.Errorf("Error while parse config file: %v", err)
		return err
	}
	db, err := repository.ConnectDB(logger, config)
	if err != nil {
		logger.Errorf("Error while connect to DB: %v", err)
		return err
	}
	repository := repository.NewRepository(db, logger)
	service := service.NewService(repository, logger)
	if err := service.AddToDB(); err != nil {
		return err
	}
	handlers := handlers.NewHandler(service)

	srv := new(test.Server)
	if err := srv.Run("8080", handlers.InitRoute()); err != nil {
		return err
	}
	return nil
}
