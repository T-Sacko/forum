package controllers

import (
	"fmt"
	m "forum/models"
	"net/http"
)

func CheckSession(w http.ResponseWriter, r *http.Request) {
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
	if loggedIn {
		w.WriteHeader(http.StatusOK)
        return
	}

	w.WriteHeader(http.StatusUnauthorized)
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
	ids := m.GetCategoriesID(categories)
	postId := m.SavePost(title, content, userId)
	m.LinkPostCategories(postId, ids)
}
