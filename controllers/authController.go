package controllers

import (
	"encoding/json"
	"fmt"

	m "forum/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {

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
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup req received")
	SessionId, err := uuid.NewV1()
	if err != nil {
		http.Error(w, "ERROR 500", http.StatusInternalServerError)
		return
	}
	user := getUser(r)
	user.SessionId = SessionId.String()
	err = user.Register()
	if err != nil {
		http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
		http.Redirect(w,r,"",http.StatusSeeOther)
		return
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: user.SessionId,
	}
	fmt.Println("man", cookie.Value)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusAccepted)
	user.Password = ""

	err = Tpl.ExecuteTemplate(w, "home.html", user)
	if err != nil {
		log.Fatal(err)
	}
	
}

var err error
var Tpl = template.Must(template.ParseGlob("/home/student/forum/templates/*.html"))

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
func getUser(r *http.Request) *m.User {
	return &m.User{Email: r.FormValue("email"), Username: r.FormValue("username"), Password: r.FormValue("password")}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	err = Tpl.ExecuteTemplate(w, "sign-in.html", nil)
	if err != nil {
		log.Fatal(err)
	}
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
