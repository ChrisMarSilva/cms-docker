package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.fasthttp
// go get -u github.com/valyala/fasthttp
// go mod tidy

// go run main.go
// go build main.go

func main() {

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			handler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	log.Printf("Server running at port: 7007")
	log.Fatal(fasthttp.ListenAndServe(":7007", requestHandler))

}

func handler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "golang api 07.Fasthttp ok\n")
}
