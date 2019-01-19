package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"

	"github.com/gorilla/sessions"
	//"time"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type LoginInformation struct {
	ID       string //`json:"ID"`
	Password string //`json:"Password"`
}

type User struct {
	ObjectType       string           //`json:"DocType"`
	LoginInformation LoginInformation //`json:"LoginInformation"`
	Name             string           //`json:"Name"`
	SocialNumber     string           //`json:"SocialNumber"`
	Location         string           //`json:"Location"`
} // user information

func getCookieName(r *http.Request) string { // cookie 이름 string으로 반환하는 부분.
	h, err := r.Header["Cookie"]
	result := ""
	if !err {
		fmt.Println("there is no cookie!")
	} else {
		for i := 0; i < len(h[0]); i++ {
			if h[0][i] == '=' {
				break
			}
			result += string(h[0][i])
		}
	}
	return result
}

func getSession(r *http.Request) { // session := store.Get 으로 가져오는 부분
	session, _ := store.Get(r, getCookieName(r))
	fmt.Println(reflect.TypeOf(session.Name()))
}

func login(w http.ResponseWriter, r *http.Request) { // 가장 문제의 메소드.
	r.ParseForm()

	id := r.Form["userID"][0]
	password := r.Form["userPW"][0]

	loginInformation := LoginInformation{
		ID:       id,
		Password: password,
	}

	/*
		아이디와 비밀번호를 담은 loginInformation을 chaincode peer로 전송.
		로그인 정보를 받아오면 쿠키 생성해야함.
	*/

	user_temp := User{ // 가라로 만든 데이터
		ObjectType:       "temp",
		LoginInformation: loginInformation,
		Name:             "kimdoyoung",
		SocialNumber:     "940901-1111111",
		Location:         "Daegu",
	}

	cookie := http.Cookie{ // cookie 세팅
		Name:     user_temp.LoginInformation.ID,
		Value:    user_temp.LoginInformation.Password,
		MaxAge:   3600, // 이 시간 초만큼 지나면 쿠키 소멸되는것 확인
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie) // 쿠키 세팅하는 방법 1
	//w.Header().Set("Set-cookie", cookie.String()) //쿠키세팅하는 방법 2

	/*session, _ := store.Get(r, getCookieName(r))
	if session.Name() == "admin" {
		session.Values["authenticated"] = true
	} else {
		session.Values["authenticated"] = false
	}
	session.Save(r, w)*/
	/*
		가장 문제되는 부분. cookie name을 가지고 session을 가지고 오는거같은데, session의 구조체 안에는 name 필드가 있다.
		근데 어떨땐 그냥 .Name으로 접근해서 가져오고, 어떨땐 .Name() 으로 가져와야한다.
		너무 혼란스러우니 다시 확인해봐야하는거같은데 뇌절온거같다. 그리고 터져버림...
		admin이 secret 페이지를 볼 수 있도록 해주는 부분이기도 하다.
	*/
}

func logout(w http.ResponseWriter, r *http.Request) { // 계정 로그아웃 시켜주는 부분. 사실 로그아웃이라기보단 admin이 secret 페이지를 못보게 만드는 부분.
	if r.Method == "GET" {
		t, _ := template.ParseFiles("test/logoutTest.html")
		t.Execute(w, nil)
	} else {
		/*session, _ := store.Get(r, getCookieName(r))

		if session.Name() == "admin" {
			session.Values["authenticated"] = false
		}
		session.Save(r, w)*/

		cookie := http.Cookie{
			Name:     getCookieName(r),
			Value:    "",
			MaxAge:   -10,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/secret", 302)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getCookieName(r))
}

func signup(w http.ResponseWriter, r *http.Request) { // 회원가입 메소드...일단 연결이 구현되면 확인
	fmt.Println("method : ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/signup.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() // 넘어온 정보들 파싱
		var obj string
		// 이제 admin의 아이디를 fix 해놓을거기 때문에, id로 사용자가 일반 사용자인지 관리자인지 판단한다.
		// 일단 민준이가 어떻게 값을 줄진 모르겠지만, objectType 을 내 임의로 string으로 넣어봤다.
		if r.Form["ID"][0] == "admin" {
			obj = "admin"
		} else {
			obj = "user"
		}

		// r.Form[data] 하면 리턴 타입이 []string임. string으로 넣어주기 위해 저렇게 붙여줘야함
		user := User{
			ObjectType: obj,
			/*ID : r.Form["ID"][0],
			Password : r.Form["Password"][0],*/

			Name:         r.Form["Name"][0],
			SocialNumber: r.Form["IDNumber"][0],
			Location:     r.Form["Region"][0],
		}

		t, _ := template.ParseFiles("templates/userVoted.html")
		t.Execute(w, user)

		/*
			회원가입 처리하면 됨
			입력받은 데이터들을 chaincode로 보내주는 부분 작성하면 됨!
		*/

		//http.Redirect(w, r, "/", 302) // 첫 화면으로 redirect
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("쿠키 이름 : ", getCookieName(r), ", 값 : ", checkCookie(r))

	if checkCookie(r) == -1 {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, nil)
	} else {
		fmt.Fprintf(w, getCookieName(r))
	}
}

func main() {
	mux := http.NewServeMux()
	//mux.Handle("/", http.FileServer(http.Dir("templates")))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/secret", secret)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/test", test)

	http.ListenAndServe(":8080", mux)
}

func secret(w http.ResponseWriter, r *http.Request) { // admin 관리자만 볼 수 있는 페이지
	/*session, _ := store.Get(r, getCookieName(r))

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "the cake is a lie!")*/

	if getCookieName(r) == "admin" {
		fmt.Fprintf(w, "관리자 페이지입니다.")
	} else if getCookieName(r) == "" {
		fmt.Fprintf(w, "다시 로그인하세요.")
	} else {
		fmt.Fprintf(w, "일반 사용자 페이지입니다.")
	}
}

func checkCookie(r *http.Request) int {
	if getCookieName(r) == "admin" {
		return 1
	} else if getCookieName(r) == "" {
		return -1
	} else {
		return 0
	}
}
