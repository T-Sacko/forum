package controllers

import (
	"database/sql"
	"fmt"
	m "forum/models"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var user *m.User

	var data *m.UserCheckResponse

	var available = false

	if r.Method == "POST" {
		switch r.URL.Path {
		case "/log-in":
			fmt.Println("log-in is happening")
			user = getUser(r)
			isUser, err := m.Check4User(user.Email, user.Password)
			user.Password = ""
			if err != nil || !isUser {
				http.Error(w, "Invalid email or password", http.StatusUnauthorized)
				return
			}
			cookie, err := CookieSetter(user)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		case "/sign-up":
			user = getUser(r)

			cookie, err := CookieSetter(user)
			if err != nil {
				http.Error(w, "500", http.StatusInternalServerError)
				return
			}

			err = user.Register()
			user.Password = ""
			if err != nil {
				http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
				return
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		case "/post":
			err := CreatePost(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", 200)
		case "/del-cookie":
			expiredCookie := http.Cookie{
				Name:    "session",
				Value:   "",
				Expires: time.Unix(0, 0),
			}
			http.SetCookie(w, &expiredCookie)
			http.Redirect(w, r, "/", http.StatusOK)
		}
	}

	_, err = r.Cookie("session")
	if err != http.ErrNoCookie {
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		user, err = m.GetUserByCookie(r)
		if err != nil {
			log.Println("Yo", err)
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		posts, err := m.GetPostsFromDB()
		if err != nil {
			log.Println(err)
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		available = true
		user = &m.User{ID: user.ID, Username: user.Username, Post: posts}
	} else {
		posts, err := m.GetPostsFromDB()
		if err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			user = &m.User{Post: posts}
		}
	}

	data = &m.UserCheckResponse{
		Available: available,
		UserInfo:  user,
	}

	err = Tpl.Execute(w, data)
	if err != nil {
		log.Println("Template execution error:", err)
	}
}
