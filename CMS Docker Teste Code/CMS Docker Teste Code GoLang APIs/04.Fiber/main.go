package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.fiber
// go get -u github.com/gofiber/fiber/v2
// go mod tidy

// go run main.go
// go build main.go

func main() {

	//app := fiber.New()

	app := fiber.New(fiber.Config{
		Prefork:      false, // doesn't run on docker even after following steps from https://docs.gofiber.io/api/fiber
		ServerHeader: "Fiber",
		AppName:      "Benchmark App",
	})

	app.Get("/", handler)

	log.Printf("Server running at port: 7004")
	log.Fatal(app.Listen(":7004"))
}

func handler(c *fiber.Ctx) error {
	return c.SendString("golang api 04.Fiber ok")
}
