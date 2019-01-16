package main

import (
	"fmt"
	"net/http"
	"html/template"
	//"encoding/json"
	//"io/ioutil"
)

func signup (w http.ResponseWriter, r *http.Request){
	fmt.Println("method : ", r.Method)
	if r.Method == "GET"{
		t, _ := template.ParseFiles("templates/signup.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() // 넘어온 정보들 파싱
		
		/*
		회원가입 처리하면 됨
		*/

		http.Redirect(w, r, "/", 302) // 첫 화면으로 redirect
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("templates")))

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)

	http.ListenAndServe("127.0.0.1:8080",mux)
}