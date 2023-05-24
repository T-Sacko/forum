package controllers

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Handle login request
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parseform", http.StatusBadRequest)
	}

	email := r.FormValue("email-address")
	username := r.FormValue("username")
	password := r.FormValue("password")
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse the ting", http.StatusBadRequest)
	}

	post := r.FormValue("post")
	category:=r.FormValue("category")



}

// Add more handler functions as needed
