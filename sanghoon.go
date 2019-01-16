package main

import (
	"net/http"
	"html/template"
//	"time"
	"fmt"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("userID: ", r.Form["userID"])
	fmt.Println("userPW: ", r.Form["userPW"])

	// get from Chaincode
	user_state := 0
	
	if user_state==1 {
		t, _ := template.ParseFiles("test/user_menu.html")
		t.Execute(w, "user")
	} else {
		t, _ := template.ParseFiles("test/admin_menu.html")
		t.Execute(w, "admin")
	}

}
