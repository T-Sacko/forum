package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
)

func CheckSession(w http.ResponseWriter, r *http.Request) {
	fmt.Println("checking session")
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
		fmt.Println("session was authorised")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func CreatePost(r *http.Request) error {
	user, err := m.GetUserByCookie(r)
	if err != nil {
		fmt.Println(err)
		return err
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["category"]
	fmt.Println(categories)
	ids := m.GetCategoriesID(categories)
	postId, err := m.SavePost(title, content, user.ID)
	if err != nil {
		return err
	}
	m.LinkPostCategories(postId, ids)
	return nil
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var Comment m.Comment

		err := json.NewDecoder(r.Body).Decode(&Comment)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		}

		user, err := m.GetUserByCookie(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Comment = m.Comment{UserID: user.ID, PostID: Comment.PostID, Comment: Comment.Comment}

		err = Comment.SaveComment()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Comment received"))
	}
}
