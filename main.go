package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("templates")))

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)


	server.ListenAndServe(":8080", nil)
}