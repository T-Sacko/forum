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
	if err != nil {
		http.Error(w, "error getting user info", http.StatusUnauthorized)
		// http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	// make comment instance
	content := r.FormValue("comment")
	if content == "" {
		http.Error(w, "can't create an empty comment", http.StatusBadRequest)
	}

	runes := []rune(content) // Convert string to slice of runes
	postID, err := strconv.Atoi(r.FormValue("postID"))
	if err != nil {
		fmt.Println("couldnt convert comment postid to int, Error", err)
		http.Error(w, "couldnt conv postID to int", http.StatusBadRequest)
	}

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
		return
	}
	var userID int
	user, _ := m.GetUserByCookie(r)
	if user != nil {
		userID = user.ID
	}
	comments, err := m.GetComments(userID, postID)
	fmt.Println(comments)
	if err != nil {
		http.Error(w,"couldnt rettieve the comments", http.StatusInternalServerError)
		fmt.Println("couldnt get comments")
		return
	}
	fmt.Println("comments length is: ", len(comments))
	json.NewEncoder(w).Encode(comments)
}

// func err5(w http.ResponseWriter){
// 	http.Error(w,"internal server err",http.StatusInternalServerError)
// }
