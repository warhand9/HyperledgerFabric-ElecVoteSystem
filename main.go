package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("templates")))

	server.ListenAndServe(":8080", nil)
}