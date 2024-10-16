package main

import (
	"log/slog"
	"os"
	"strconv"

	"host/internal/api"
	"host/internal/api/controllers"
	"host/internal/domain"
	storagePostgres "host/internal/storage-postgres"
	"host/internal/storage-postgres/repositories"
	"host/internal/utils"
)

func main() {
	logger := slog.Default()
	slog.SetDefault(logger)

	config := utils.ParseConfig(map[string]string{
		"PORT":        "8080",
		"DB_USER":     "",
		"DB_PORT":     "",
		"DB_HOST":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
	})

	dbPort, err := strconv.Atoi(*config["DB_PORT"])
	if err != nil {
		logger.Error("Error parsing DB_PORT to int")
		os.Exit(1)
	}

	dbConfig := storagePostgres.DBConfig{
		Port:     dbPort,
		Host:     *config["DB_HOST"],
		User:     *config["DB_USER"],
		Password: *config["DB_PASSWORD"],
		DBName:   *config["DB_NAME"],
	}

	db, err := storagePostgres.InitDB(dbConfig)
	if err != nil {
		logger.Error("Failed to connect to database:", err)
		os.Exit(1)
	}

	repo := repositories.NewUserRepository(db)
	service := domain.NewUserService(repo)
	controller := controllers.NewUserController(service)

	port := config["PORT"]
	server := api.NewServer(*port, *controller)

	if err := server.Init(); err != nil {
		logger.Error("Failed to start server:", err)
		os.Exit(1)
	}
}
