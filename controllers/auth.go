package controllers

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var err error

var Tpl *template.Template

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println(username)
		HashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		err = models.InsertDB(username, email, string(HashedPass))
		if err != nil {
			log.Fatal(err)
		}
	}
	err = Tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	err = Tpl.ExecuteTemplate(w, "sign-in.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	err = Tpl.ExecuteTemplate(w, "create_post.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func UsernameCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	available, err := models.IsUsernameAvailable(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := models.UsernameCheckResponse{
		Available: available,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

/*
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
*/
