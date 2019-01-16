package main

import (
	"net/http"
	"html/template"
	"time"
	"fmt"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("templates")))

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)


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
}