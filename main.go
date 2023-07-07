package main

import (
	"log"

	"github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/cmd/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func main() {
	route := app.StartApp()

	log.Println("Starting the application...")
	route.Start(":8080")
}
