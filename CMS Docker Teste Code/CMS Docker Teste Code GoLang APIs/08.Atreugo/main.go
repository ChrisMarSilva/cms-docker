package main

import (
	"log"

	"github.com/savsgio/atreugo/v11"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.atreugo
// go get -u github.com/savsgio/atreugo/v11
// go mod tidy

// go run main.go
// go build main.go

func main() {

	server := atreugo.New(atreugo.Config{Addr: "0.0.0.0:7008"})
	server.GET("/", handler)

	log.Printf("Server running at port: 7008")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func handler(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("golang api 08.Atreugo ok")
}
