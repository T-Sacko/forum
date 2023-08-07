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

	fmt.Println("commentddddddddddddddddd", r.FormValue("comment"))
	comment := &m.Comment{
		Username: user.Username,
		Content:  r.FormValue("comment"),
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
