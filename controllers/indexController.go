package controllers

import (
	"fmt"
	"log"
	"net/http"
	m"forum/models"
	"github.com/gofrs/uuid"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		switch r.URL.Path {
		case "/log-in":
			user := getUser(r)
			isUser, err := m.Check4User(user.Email, user.Password)
			fmt.Println(isUser, user.Password)
			if err != nil {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}
			// Set the session ID and create a cookie
			SessionId, err := uuid.NewV1()
			if err != nil {
				http.Error(w, "ERROR 500", http.StatusInternalServerError)
				return
			}
			m.SetSessionId(user.Email, SessionId.String())
			cookie := &http.Cookie{
				Name:  "session",
				Value: SessionId.String(),
				// Session cookie (valid until browser is closed)
			}
			http.SetCookie(w, cookie)
			// Redirect to the home page or a dashboard page
			http.Redirect(w, r, "/", http.StatusFound)
		default:
			
			SessionId, err := uuid.NewV1()
			if err != nil {
				http.Error(w, "ERROR 500", http.StatusInternalServerError)
				return
			}
			user := getUser(r)
			err = user.Register()
			if err != nil {
				http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
				return
			}
			user.SessionId = SessionId.String()
			cookie := &http.Cookie{
				Name:  "session",
				Value: user.SessionId,
			}
			fmt.Println("man",cookie.Value)
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusAccepted)
			user.Password = ""
		
			err = Tpl.ExecuteTemplate(w, "home.html", user)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	Tpl.ExecuteTemplate(w, "home.html", nil)
}
