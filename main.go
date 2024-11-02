package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString:= os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not defined")
	}
	fmt.Println("Server running on port", portString)
}
