package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.gin
// go get -u github.com/gin-gonic/gin
// go mod tidy

// go run main.go
// go build main.go

func main() {

	router := gin.Default()
	router.GET("/", handler)

	log.Printf("Server running at port: 7003")
	log.Fatal(router.Run(":7003"))
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "golang api 03.Gin ok"})
}
