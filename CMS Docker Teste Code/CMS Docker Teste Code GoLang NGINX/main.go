package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.nginx
// go mod tidy

// go run main.go
// go build main.go

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	name, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Oops: %v\n", err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Fprintf(w, "Oops: %v\n", err)
		return
	}

	for _, a := range addrs {
		fmt.Fprintf(w, a, "\n")
	}

	fmt.Fprintf(w, "PID: ", os.Getpid(), "\n")
	fmt.Fprintf(w, "RemoteAddr: ", r.RemoteAddr, "\n")

	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)

		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
		return
	}

	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr
	// Header.Get is case-insensitive
	forward := r.Header.Get("X-Forwarded-For")

	fmt.Fprintf(w, "<p>IP: %s</p>", ip, "\n")
	fmt.Fprintf(w, "<p>Port: %s</p>", port, "\n")
	fmt.Fprintf(w, "<p>Forwarded for: %s</p>", forward, "\n")

	fmt.Fprintf(w, "Hello World", "\n")
	// w.Write([]byte("ok 1"))
}
