package main

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
	"fmt"
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