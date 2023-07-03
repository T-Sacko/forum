package controllers

import (
	m "forum/models"
	"log"
	"net/http"
)

var user *m.User

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		switch r.URL.Path {
		case "/log-in":
			user = getUser(r)
			isUser, err := m.Check4User(user.Email, user.Password)
			if err != nil || !isUser {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}
			cookie, err := CookieSetter(user)
			if err != nil {
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		case "/sign-up":
			user = getUser(r)
			err = user.Register()
			if err != nil {
				http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
				return
			}
			cookie, err := CookieSetter(user)
			if err != nil {
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusAccepted)
			user.Password = ""
		case "/post":
			CreatePost(w, r)
			
			http.Redirect(w, r, "/", 200)
		}
	}
	err := Tpl.ExecuteTemplate(w, "home.html", user)
	if err != nil {
		log.Fatal(err)
	}
}
