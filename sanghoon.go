package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type User_response struct {
	Name     string
	Region   string
	Response int
}

// get all vote information
type Vote struct {
	Name           string
	StartDate      string
	EndDate        string
	User_responses []User_response
}
type EveryVote struct {
	Vote_Results []Vote
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("mainIndex")))

	mux.HandleFunc("/login", login)
	//	mux.HandleFunc("/signup", signup)

	mux.HandleFunc("/view_result", view_result)
	mux.HandleFunc("/view_vote_result", view_vote_result)

	mux.HandleFunc("/enroll_vote", enroll_vote)
	mux.HandleFunc("/membership_manage", membership_manage)

	http.ListenAndServe(":8080", mux)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("userID: ", r.Form["userID"])
	fmt.Println("userPW: ", r.Form["userPW"])

	fmt.Println("sex1: ", r.Form["voteID"])

	// get from Chaincode
	user_state := 0

	if user_state == 1 {
		t, _ := template.ParseFiles("test/user_menu.html")
		t.Execute(w, "user")
	} else {
		t, _ := template.ParseFiles("test/admin_menu.html")
		t.Execute(w, "admin")
	}

}

func view_result(w http.ResponseWriter, r *http.Request) {
	vote_result := EveryVote{
		Vote_Results: []Vote{
			{
				Name:      "Daetongyeong_2019_vote",
				StartDate: "2019/01/17",
				EndDate:   "2019/01/17",
				User_responses: []User_response{
					{Name: "LSH", Region: "Daemeong", Response: 1},
					{Name: "ABC", Region: "Daemeong", Response: 2},
					{Name: "DEF", Region: "Suseong", Response: 3},
				},
			},
			{
				Name:      "Gukhaeuione_2019_vote",
				StartDate: "2019/01/16",
				EndDate:   "2019/01/16",
				User_responses: []User_response{
					{Name: "ZZZ", Region: "Suseong", Response: 2},
					{Name: "QQQ", Region: "Daemeong", Response: 1},
					{Name: "AAA", Region: "Suseong", Response: 3},
				},
			},
			{
				Name:      "Daetongyeong_2018_vote",
				StartDate: "2018/01/17",
				EndDate:   "2018/01/17",
				User_responses: []User_response{
					{Name: "QEW", Region: "Suseong", Response: 1},
					{Name: "ZXC", Region: "Daemeong", Response: 1},
					{Name: "HJK", Region: "Suseong", Response: 1},
				},
			},
			{
				Name:      "Gukhaeuione_2018_vote",
				StartDate: "2018/01/16",
				EndDate:   "2018/01/16",
				User_responses: []User_response{
					{Name: "BBB", Region: "Daemeong", Response: 1},
					{Name: "WEE", Region: "Suseong", Response: 2},
					{Name: "PPP", Region: "Suseong", Response: 3},
				},
			},
		},
	}

	jsonBytes, _ := json.Marshal(vote_result)
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	t, _ := template.ParseFiles("test/view_result.html")
	t.Execute(w, jsonString)
}

func view_vote_result(w http.ResponseWriter, r *http.Request) {
	fmt.Println("vote ID: ", r.Form["vote_id"])

	vote_result := Vote{
		Name:      "Daetongyeong_2019_vote",
		StartDate: "2019/01/17",
		EndDate:   "2019/01/17",
		User_responses: []User_response{
			{Name: "LSH", Region: "Daemeong", Response: 1},
			{Name: "ABC", Region: "Daemeong", Response: 2},
			{Name: "DEF", Region: "Suseong", Response: 3},
		},
	}

	jsonBytes, _ := json.Marshal(vote_result)
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	t, _ := template.ParseFiles("test/view_vote_result.html")
	t.Execute(w, jsonString)
}

func enroll_vote(w http.ResponseWriter, r *http.Request) {

}

func membership_manage(w http.ResponseWriter, r *http.Request) {

}
