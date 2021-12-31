package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //signup.htmlファイルを出力
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		}else{
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" { //データが入力された場合の処理
		err := r.ParseForm() //解析する
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{ //ユーザーのstructの各フィールドの値として受け取る
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil { //新規ユーザー登録をする
			log.Println(err)
		}
		http.Redirect(w, r, "/", 302) //登録に成功したらトップページにリダイレクトする
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "login")
			}else{
				http.Redirect(w, r, "/todos", 302)
			}
		}
	

func auhtenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email")) //画面から情報を取得
	if err != nil {
		log.Println(err)
		http.Redirect(w,r, "/login", 302)
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {//暗号化する
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		} 
		//
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	}else {
		http.Redirect(w,r,"/login",302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DelateSessionByUUID()
	}
	http.Redirect(w,r, "/login", 302)
}


