package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
)

type SessionStatusResponse struct {
	LoggedIn bool `json:"loggedIn"`
}

func CheckSession(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ajax active session request received")
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "unathorized to post", http.StatusUnauthorized)
	}

	sessionId:=cookie.Value

	loggedIn, err := m.SessionIsActive(sessionId)

	if err != nil {
		fmt.Println(err, "problem with posting because of cookie")
	}

	response := SessionStatusResponse{
		LoggedIn: loggedIn,
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("cant marshal the response into json")

	}

}

 func getUser


func CreatePost(w http.ResponseWriter, r *http.Request) {
	user := getUser()
}
