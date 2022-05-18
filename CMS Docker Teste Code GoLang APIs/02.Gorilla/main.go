package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.gorilla
// go get -u github.com/gorilla/mux
// go mod tidy

// go run main.go
// go build main.go

func main() {
	router := mux.NewRouter() //.StrictSlash(true)
	router.HandleFunc("/", handler)
	log.Printf("Server running at port: 7002")
	log.Fatal(http.ListenAndServe(":7002", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("golang api 02.Gorilla ok "))
}
