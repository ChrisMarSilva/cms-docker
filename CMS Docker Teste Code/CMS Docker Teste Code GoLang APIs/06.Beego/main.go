package main

import (
	"log"

	"github.com/beego/beego/v2/server/web"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.beego
// go get -u github.com/beego/beego/v2
// go mod tidy

// go run main.go
// go build main.go

func main() {

	ctrl := &MainController{}

	//web.Get("/", ctrl)
	web.Router("/", ctrl)

	log.Printf("Server running at port: 7006")
	web.Run(":7006")

}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) Get() {
	ctrl.Ctx.WriteString("golang api 06.Beego ok")
}
