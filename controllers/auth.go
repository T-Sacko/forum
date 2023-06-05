package controllers

import (
	"fmt"
	m "forum/models"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Failed to parse the ting", http.StatusBadRequest)
// 	}

// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	// Handle login request
// }

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "Can't parse da html", http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to execute HTML template", http.StatusInternalServerError)
		return
	}

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// if method is POST then handle user registration

	if r.Method == "POST" {
		email := r.FormValue("email-address")
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := m.User{
			Email:    email,
			Username: username,
			Password: password,
		}

		usernameAvailable, _ := m.IsUsernameAvailable(username)
		emailAvailable, _ := m.IsEmailAvailable(email)

		if usernameAvailable && emailAvailable {

			err := user.Save()
			if err != nil {
				fmt.Println("what")
				http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
				return
			}
			sessionID := uuid.New().String()
			m.SetSessionId(sessionID)
			cookie := &http.Cookie{
				Name:   "session",
				Value:  sessionID,
			}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusFound)
		}

		// if not then Parse register.html

	} else {
		tmpl, err := template.ParseFiles("register.html")
		if err != nil {
			http.Error(w, "Can't parse da html", http.StatusInternalServerError)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Failed to execute HTML template", http.StatusInternalServerError)
			return
		}
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := m.CheckUserCredentials(email, password)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Set the session ID and create a cookie
		sessionID := uuid.New().String()
		user.SetSessionID(sessionID)
		cookie := &http.Cookie{
			Name:   "session",
			Value:  sessionID,
			MaxAge: -1, // Session cookie (valid until browser is closed)
		}
		http.SetCookie(w, cookie)

		// Redirect to the home page or a dashboard page
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// Display the login form
		tmpl, err := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, "Can't parse the HTML", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Failed to execute HTML template", http.StatusInternalServerError)
			return
		}
	}
}


// func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "failed to parse the ting", http.StatusBadRequest)
// 	}

// 	post := r.FormValue("post")
// 	category := r.FormValue("category")

// }

// Add more handler functions as needed
