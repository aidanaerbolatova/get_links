package app

import (
	"test/internal/repository"
	"test/internal/service"
	test "test/internal/transport/httpserver"
	"test/internal/transport/httpserver/handlers"
)

func Run() error {
	db, err := repository.ConnectDB()
	if err != nil {
		return err
	}
	repository := repository.NewRepository(db)
	data, err := service.ConvertJson("links.json")
	if err != nil {
		return err
	}
	service := service.NewService(repository)
	if err := service.AddToDB(data); err != nil {
		return err
	}
	handlers := handlers.NewHandler(service)

	srv := new(test.Server)
	if err := srv.Run("8080", handlers.InitRoute()); err != nil {
		return err
	}
	return nil
}
