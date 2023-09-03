package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	m "forum/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

var err error
var Tpl = template.Must(template.ParseGlob("templates/*.html"))

// type loggedIn struct {
// 	LoggedIn bool `json:"loggedIn"`
// }

func Session(w http.ResponseWriter, r *http.Request) {
	var seshData struct {
		Status bool `json:"status"`
	}
	_, err := m.GetUserByCookie(r)
	if err != nil {
		json.NewEncoder(w).Encode(seshData)
		return
	}
	seshData.Status = true
	json.NewEncoder(w).Encode(seshData)

}

func cookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", err
	}
	seshID := cookie.Value
	return seshID, nil

}

// ////////////////////////////////////////////////
func SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("log out req received")
	seshID, err := cookie(r)
	if err != nil {
		fmt.Println("next tings on signout:", err)
	}

	err = m.DeleteCookie(seshID)
	if err != nil {
		fmt.Println("cant log out:", err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_cookie_name", // replace with the name of your session cookie
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0), // Setting it to the zero value makes it expired
		MaxAge:  -1,              // MaxAge<0 means delete cookie now
	})
	// http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// ////////////////////////////////////////////////////
func Login(w http.ResponseWriter, r *http.Request) {

	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// user1 := getUser(r)
	isUser, err := m.Check4User(user.Email, user.Password)
	fmt.Println(isUser, user.Password)
	if err != nil {
		fmt.Println(isUser, user.Password, "not nil")

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Set the session ID and create a cookie
	SessionId, err := uuid.NewV1()
	if err != nil {
		http.Error(w, "Couldn't create sesh id", http.StatusInternalServerError)
		return
	}
	m.SetSessionId(user.Email, SessionId.String())
	cookie := &http.Cookie{
		Name:  "session",
		Value: SessionId.String(),
		// Session cookie (valid until browser is closed)
	}

	http.SetCookie(w, cookie)
	fmt.Println("yessssss")
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
		http.Redirect(w, r, "", http.StatusSeeOther)
		return
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: user.SessionId,
	}
	fmt.Println("man", cookie.Value)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)

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
	fmt.Println("the username", username, "is available", available)
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
