package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// |->
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

// <-|

//
type VoteResult struct {
	Name      string
	StartDate string
	EndDate   string
}

//

type User struct {
	Name   string
	Unum   string
	ID     string
	PW     string
	Region string
}
type UserCount struct {
	Region string
	Count  int
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("mainIndex")))

	//mux.HandleFunc("/", indexFunc)
	mux.HandleFunc("/login", login)
	//	mux.HandleFunc("/signup", signup)

	mux.HandleFunc("/view_result", view_result)
	mux.HandleFunc("/view_vote_result", view_vote_result)

	mux.HandleFunc("/enroll_vote", enroll_vote)
	mux.HandleFunc("/membership_manage", membership_manage)

	http.ListenAndServe(":8084", mux)
}

/*func indexFunc(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("mainIndex/index.html")
	t.Execute(w, "test")
}*/

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("userID: ", r.Form["userID"])
	fmt.Println("userPW: ", r.Form["userPW"])

	fmt.Println("hidden test: ", r.Form["voteID"])

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

// usercounts function
func getUserCount(jsonString string) string {
	testInfo := []User{
		{
			Name:   "AAA",
			Unum:   "1-1",
			ID:     "aaa",
			PW:     "1234",
			Region: "Daemeong",
		},
		{
			Name:   "BBB",
			Unum:   "2-2",
			ID:     "bbb",
			PW:     "1234",
			Region: "Suseong",
		},
		{
			Name:   "CCC",
			Unum:   "3-3",
			ID:     "ccc",
			PW:     "1234",
			Region: "Daemeong",
		},
		{
			Name:   "DDD",
			Unum:   "4-4",
			ID:     "ddd",
			PW:     "1234",
			Region: "Daemeong",
		},
		{
			Name:   "EEE",
			Unum:   "5-5",
			ID:     "eee",
			PW:     "1234",
			Region: "Suseong",
		},
		{
			Name:   "FFF",
			Unum:   "6-6",
			ID:     "fff",
			PW:     "1234",
			Region: "Suseong",
		},
	}

	jsonBytes, _ := json.Marshal(testInfo)
	jsonString = string(jsonBytes)

	///////

	var userlists []User
	json.Unmarshal([]byte(jsonString), &userlists)
	//fmt.Println("unmarshal test", len(userlists))

	//userCountInfo := make([]UserCount, 0)
	var userCountInfo []UserCount
	// type UserCount struct {Region string Count int}

	for i := 0; i < len(userlists); i++ {
		temp_region := userlists[i].Region
		if regionNotExists(temp_region, userCountInfo) {
			var new_user_count UserCount = UserCount{temp_region, 0}
			userCountInfo = append(userCountInfo, new_user_count)
		}

		indx := getIndex(temp_region, userCountInfo)
		userCountInfo[indx].Count++
	}

	//
	resultBytes, _ := json.Marshal(userCountInfo)
	resultString := string(resultBytes)

	return resultString
}

func regionNotExists(temp_region string, userCountInfo []UserCount) bool {
	for i := 0; i < len(userCountInfo); i++ {
		if userCountInfo[i].Region == temp_region {
			return false
		}
	}

	return true
}
func getIndex(temp_region string, userCountInfo []UserCount) int {
	for i := 0; i < len(userCountInfo); i++ {
		if userCountInfo[i].Region == temp_region {
			return i
		}
	}

	return -1
}

// usercounts function end

// or doThis(jsonStringData String) (string)
func doThis(jsonByteData []byte) string {

	var data []Vote
	// json.Unmarshal([]byte(jsonString), &data)
	json.Unmarshal(jsonByteData, &data)

	//var result []VoteResult
	result := make([]VoteResult, len(data)) //,,INSTEAD,, USE append() functio

	for i := 0; i < len(data); i++ {
		var temp VoteResult
		temp.Name = data[i].Name
		temp.StartDate = data[i].StartDate
		temp.EndDate = data[i].EndDate

		result[i] = temp
	}

	resultBytes, _ := json.Marshal(result)
	resultString := string(resultBytes)
	//fmt.Println(resultString)

	return resultString
}
func view_result(w http.ResponseWriter, r *http.Request) {

	//
	/*vote_result := EveryVote{
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
	}*/
	vote_result := []Vote{
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
	}

	// example data received from chaincode(JsonString) -> presenting as Go structure

	jsonBytes, _ := json.Marshal(vote_result)
	//jsonString := string(jsonBytes)
	//fmt.Println(jsonString)
	// example JsonString supposed as being received from chaincode

	// parse this JsonString & get requiremenet data and make Result JsonString & send to tempaltes
	resultString := doThis(jsonBytes) // or doThis(jsonString) depend on whether data format chaincode returns is jsonBytes or jsonString
	fmt.Println("RESULT STRING :: " + resultString)

	// get user count information
	userCountInfos := getUserCount("jsonstring received from chaincode")
	//fmt.Println("USERCOUNTINFOS : ", userCountInfos)
	resultString = resultString + "splitSeparator" + userCountInfos
	fmt.Println("VoteReuslt + UserCountInfos : ", resultString)

	t, _ := template.ParseFiles("test/view_result.html")
	//t.Execute(w, jsonString)
	t.Execute(w, resultString)
}

func view_vote_result(w http.ResponseWriter, r *http.Request) {
	fmt.Println("voteID: ", r.FormValue("voteID"))

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

type enrollInfo struct {
	VoteName string
	EndDate  string
	Articles []string
}

func enroll_vote(w http.ResponseWriter, r *http.Request) {
	var enrollRequest enrollInfo
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &enrollRequest)
	fmt.Println(enrollRequest.VoteName, enrollRequest.EndDate, enrollRequest.Articles)
	w.Write([]byte("success!!!!!"))
}

func membership_manage(w http.ResponseWriter, r *http.Request) {
	testInfo := []User{
		{
			Name:   "AAA",
			Unum:   "1-1",
			ID:     "aaa",
			PW:     "1234",
			Region: "Daemeong",
		},
		{
			Name:   "BBB",
			Unum:   "2-2",
			ID:     "bbb",
			PW:     "1234",
			Region: "Suseong",
		},
		{
			Name:   "CCC",
			Unum:   "3-3",
			ID:     "ccc",
			PW:     "1234",
			Region: "Daemeong",
		},
		{
			Name:   "DDD",
			Unum:   "4-4",
			ID:     "ddd",
			PW:     "1234",
			Region: "Daemeong",
		},
		{
			Name:   "EEE",
			Unum:   "5-5",
			ID:     "eee",
			PW:     "1234",
			Region: "Suseong",
		},
		{
			Name:   "FFF",
			Unum:   "6-6",
			ID:     "fff",
			PW:     "1234",
			Region: "Suseong",
		},
	}

	jsonBytes, _ := json.Marshal(testInfo)
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	t, _ := template.ParseFiles("test/membership_manage.html")
	t.Execute(w, jsonString)

}
