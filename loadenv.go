package main

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}
}
