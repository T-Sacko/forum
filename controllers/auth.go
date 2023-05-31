package controllers

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"html/template"
	"log"
	"net/http"
)

var err error

var Tpl *template.Template

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("DONE")
		switch r.URL.Path {
		case "/log-in":
			auth, check := SigningIn(w, r)
			if auth == "Authenticated" {
				fmt.Println(auth)
				err = Tpl.ExecuteTemplate(w, "home.html", check.UserInfo)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				http.Error(w, auth, http.StatusNoContent)
			}
		default:
			err = getUser(r).Save()
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(w, r, "/", http.StatusAccepted)
		}
	}
	err = Tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getUser(r *http.Request) *models.User {
	return &models.User{Email:r.FormValue("email"), Username:r.FormValue("username"), Password: r.FormValue("password")}
}


func SigningIn(w http.ResponseWriter, r *http.Request) (string, models.UserCheckResponse) {
	check, err := getUser(r).LogIn()
	if err != nil {
		log.Fatal(err)
	}
	if check.Available {
		return "Authenticated", check
	}
	return "Unauthenticated", models.UserCheckResponse{}
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

	response := models.UserCheckResponse{
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

func EmailCheck(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	available, err := models.IsEmailAvailable(email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := models.UserCheckResponse{
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
