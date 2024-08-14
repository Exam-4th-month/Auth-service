package main

import (
	"log"
	"log/slog"
	"os"

	"auth-service/api"
	"auth-service/internal/items/config"

	"auth-service/internal/items/service"
	"auth-service/internal/items/storage"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	db, err := storage.ConnectDB(config)
	if err != nil {
		logger.Error("error while connecting postgres:", slog.String("err:", err.Error()))
	}

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	service := service.New(storage.New(
		db,
		sqrl,
		config,
		logger,
	), logger)

	api := api.New(service)

	log.Fatalln(api.RUN(config, service))

}
