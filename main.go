package main

import (
	"net/http"
	//"fmt"
)

func main_temp() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("templates")))

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)

	http.ListenAndServe(":8080", mux)
}
