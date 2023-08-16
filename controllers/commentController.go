package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
	"strconv"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {

	// get user details
	user, err := m.GetUserByCookie(r)
	fmt.Print("\nreceived comment\n\n")
	if err != nil {
		// http.Error(w, "error getting user info", http.StatusUnauthorized)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	// make comment instance
	postID, err := strconv.Atoi(r.FormValue("postID"))
	if err != nil {
		fmt.Println("couldnt convert comment postid to int, Error", err)
		http.Error(w, "couldnt conv postID to int", http.StatusBadRequest)
	}

	content := r.FormValue("comment")
	runes := []rune(content) // Convert string to slice of runes

	// Check if the last character is not an ASCII character, and if not, remove it
	for len(runes) > 0 && (runes[len(runes)-1] < 32 || runes[len(runes)-1] > 127) {
		runes = runes[:len(runes)-1]
	}
	content = string(runes)
	fmt.Println(content, "reee")
	comment := &m.Comment{
		Username: user.Username,
		Content:  content,
		PostID:   postID,
		UserID:   user.ID,
	}

	ID, err := comment.Save()

	comment.ID = int(ID)

	if err != nil {
		http.Error(w, "cant save comment", http.StatusInternalServerError)
		fmt.Println("couldnt save comment, Error:", err)
		return
	}

	json.NewEncoder(w).Encode(comment)

}

func GetComments(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("postID")
	postID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("postID is invalid to get comments")
	}
	user, err := m.GetUserByCookie(r)
	if err != nil {
		user = nil
	}
	comments, err := m.GetComments(user.ID, postID)
	fmt.Println(comments)
	if err != nil {
		fmt.Println("couldnt get comments")
	}
	fmt.Println("comments length is: ", len(comments))
	json.NewEncoder(w).Encode(comments)
}
