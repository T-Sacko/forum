package controllers

import (
	"fmt"
	m "forum/models"
	"html/template"
	"net/http"
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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email-address")
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := m.User{
		Email:    email,
		Username: username,
		Password: password,
	}

	err := user.Save()
	if err != nil {
		fmt.Println("what")
		http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)

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
