package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server.ListenAndServe(":8080", nil)
}