package main

import (
	"fmt"
	"net/http"
	"html/template"
<<<<<<< HEAD
	"time"
	"fmt"
=======
	//"encoding/json"
	//"io/ioutil"
>>>>>>> 9c1b9402708e8db7b4e4d59ecbb7095cb240feb0
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

<<<<<<< HEAD

	server.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("userID: ", r.Form["userID"])
	fmt.Println("userPW: ", r.Form["userPW"])

	// get from Chaincode
	user_state := 0
	
	if user_state==1 {
		t, _ := template.ParseFiles("user_menu.html")
	}
	else {
		t, _ := template.ParseFiles("admin_menu.html")
	}
	t.Execute(w, nil)
*/
=======
	http.ListenAndServe("127.0.0.1:8080",mux)
>>>>>>> 9c1b9402708e8db7b4e4d59ecbb7095cb240feb0
}