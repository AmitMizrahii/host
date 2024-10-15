package main

import (
	"log"
	"os"
	"strconv"

	"host/api"
	"host/dal"
	"host/domain"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USRE")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error parsing DB_PORT to int")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConfig := dal.DBConfig{
		Port:     dbPort,
		Host:     dbHost,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}
	db, err := dal.InitDB(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := dal.NewUserRepository(db)
	service := domain.NewUserService(repo)
	controller := api.NewUserController(service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := api.NewServer(port, *controller)

	if err := server.Init(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
