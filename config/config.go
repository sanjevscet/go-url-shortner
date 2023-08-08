package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDBConnectionString() string {
	//load env from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error in loading .env file %v", err)
	}

	dbUser := os.Getenv("DB_USER")

	log.Println(dbUser)
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	return connectionString
}
