package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("request received")
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "Hello World"
	}
	fmt.Fprintf(w, "%s!", greeting)
}

func main() {
	log.Print("listening on 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

