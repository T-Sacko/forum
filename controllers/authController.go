package controllers

import (
	"encoding/json"

	m "forum/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

var err error
var Tpl = template.Must(template.ParseGlob("/home/student/forum/templates/*.html"))

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
func getUser(r *http.Request) *m.User {
	return &m.User{Email: r.FormValue("email"), Username: r.FormValue("username"), Password: r.FormValue("password")}
}

func CookieSetter(user *m.User) (*http.Cookie, error) {
	SessionId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	user.SessionId = SessionId.String()
	cookie := &http.Cookie{
		Name:  "session",
		Value: user.SessionId,
	}
	m.SetSessionId(user.Email, user.SessionId)
	return cookie, nil
}

// ----------------------------------------------------ajax

func UsernameCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	available, err := m.IsUsernameAvailable(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response := m.UserCheckResponse{
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
	available, err := m.IsEmailAvailable(email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response := m.UserCheckResponse{
		Available: available,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("can't encode email check response into json", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
