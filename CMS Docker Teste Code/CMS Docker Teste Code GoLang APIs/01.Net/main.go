package main

import (
	"log"
	"net/http"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.net
// go get -u github.com/gorilla/mux
// go mod tidy

// go run main.go
// go build main.go

func main() {
	http.HandleFunc("/", handler)
	log.Printf("Server running at port: 7001")
	log.Fatal(http.ListenAndServe(":7001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("golang api 01.Net ok "))
}
