package controllers

import (
	"fmt"
	m "forum/models"
	"log"
	"net/http"
	"time"
)

var user *m.User

var data *m.UserCheckResponse

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		switch r.URL.Path {
		case "/log-in":
			user = getUser(r)
			isUser, err := m.Check4User(user.Email, user.Password)
			user.Password = ""
			if err != nil || !isUser {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}
			cookie, err := CookieSetter(user)
			if err != nil {
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}
			data.Available = true
			http.SetCookie(w, cookie)
			fmt.Println("YO", data.Available)
			http.Redirect(w, r, "/", http.StatusFound)
		case "/sign-up":
			user = getUser(r)
			err = user.Register()
			user.Password = ""
			if err != nil {
				http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
				return
			}
			cookie, err := CookieSetter(user)
			if err != nil {
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}

			fmt.Println("Here", data.Available)
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusAccepted)
		case "/post":
			file, header, err := r.FormFile("image")
			if err != nil {
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}
			fmt.Println(header)
			defer file.Close()

			// CreatePost(w, r)

			http.Redirect(w, r, "/", 200)
		case "/del-cookie":
			fmt.Println("yo")
			expiredCookie := http.Cookie{
				Name: "session",
				Value: "",
				Expires: time.Unix(0, 0),
			}
			http.SetCookie(w, &expiredCookie)
			w.WriteHeader(http.StatusOK)
		}
	}

	data = &m.UserCheckResponse{
		UserInfo: user,
	}
	if data.UserInfo != nil {
		data.Available = true
	}
	err := Tpl.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		log.Fatal(err)
	}
}
