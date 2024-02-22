package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.fiber
// go get -u github.com/gofiber/fiber/v3
// go install github.com/cosmtrek/air@latest
// go mod tidy

// go run main.go
// air

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7004"
	}

	log.Printf("Server running at port: " + port)

	app := fiber.New()
	app.Get("/", handler)
	log.Fatal(app.Listen(":" + port))
}

func handler(c fiber.Ctx) error {
	return c.SendString("Golang api 04.Fiber OK")
}
