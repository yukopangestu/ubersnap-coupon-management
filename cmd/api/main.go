package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/username/go-webapp/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloHandler)

	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
