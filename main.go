package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	a := App{}
	a.Initialize(os.Getenv("APP_DATABASE"))

	a.Run(":8000")
}
