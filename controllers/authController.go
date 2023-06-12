package controllers

import (
	"encoding/json"
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
func getUser(r *http.Request) *models.User {
	return &models.User{Email: r.FormValue("email"), Username: r.FormValue("username"), Password: r.FormValue("password")}
	return &models.User{Email: r.FormValue("email"), Username: r.FormValue("username"), Password: r.FormValue("password")}
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
