package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.fiber
// go get -u github.com/gofiber/fiber/v3
// go get -u github.com/cosmtrek/air
// go mod tidy

// air -c .air.toml
// air server --port 8080

// go run main.go

func main() {
	app := fiber.New()

	app.Get("/", handler)

	log.Printf("Server running at port: 7004")
	log.Fatal(app.Listen(":7004"))
}

func handler(c fiber.Ctx) error {
	return c.SendString("golang api 04.Fiber ok")
}
