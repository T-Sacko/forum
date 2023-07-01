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
		fmt.Println("There is no cookie")
		return
	}

	sessionId := cookie.Value

	loggedIn, err := m.SessionIsActive(sessionId)

	if err != nil {
		http.Error(w, "unathorized: invalid sesh id", http.StatusUnauthorized)
		fmt.Println(err, "invalid sesh id")
		return
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("making post began")
	userId, err := m.GetUserByCookie(r)
	if err != nil {
		http.Error(w, "user has no cookie", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["category"]
	fmt.Println(categories)		
	ids:= m.GetCategoriesID(categories)
	postId := m.SavePost(title, content,  userId)
	m.LinkPostCategories(postId,ids)
}
