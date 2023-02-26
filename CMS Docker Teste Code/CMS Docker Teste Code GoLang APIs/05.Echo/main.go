package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.echo
// go get -u github.com/labstack/echo/v4
// go mod tidy

// go run main.go
// go build main.go

func main() {

	e := echo.New()
	e.GET("/", handler)

	log.Printf("Server running at port: 7005")
	e.Logger.Fatal(e.Start(":7005"))
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "golang api 05.Echo ok")
}
