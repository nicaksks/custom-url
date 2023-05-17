package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}

func Port() string {
	Config()
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Printf("\nPort is missing in .env! \nFor security, your port has been set to 8080\n\n")
		port = "8080"
	}
	return ":" + port
}

func Database() string {
	Config()
	database := os.Getenv("DATABASE")
	if database == "" {
		log.Fatal("Database URL is missing in .env!")
	}
	return database
}

func Domain() string {
	Config()
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("Domain URL is missing in .env! For security, your domain has been set to https://localhost/")
		domain = "https://localhost/"
	}
	return domain
}
