package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
	"strconv"
)

func HandlePostLikes(w http.ResponseWriter, r *http.Request) {
	// sessionId, _ := checkCookie(w, r)
	// userId, session, err1 := m.SessionIsActive(sessionId)
	user, err := m.GetUserByCookie(r)
	if err != nil {
		http.Error(w, "unauthorized, u aint logged in", http.StatusUnauthorized)
		fmt.Println("post.controller.go Error func HandlePostAction: ", err)
		return
	}

	postActionReqData := m.PostActionReq{}
	json.NewDecoder(r.Body).Decode(&postActionReqData)
	fmt.Println("likeActionData una", postActionReqData)
	postId, _ := strconv.Atoi(postActionReqData.PostId)
	action := postActionReqData.Action

	switch action {
	case "like":
		m.SaveLike(postId, user.ID)
	case "unlike":
		m.RemoveLike(postId, user.ID)
	case "dislike":
		m.SaveDislike(postId, user.ID)
	case "removeDislike":
		m.RemoveDislike(postId, user.ID)
	}

}

// type CommentLike struct {
// 	CommentId int
// 	action string
// }

// var CommentQueue []CommentLike

func HandleCommentLikes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at least")
	user, err := m.GetUserByCookie(r)
	// if the user aint logged take them to the home page
	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	action := r.URL.Query().Get("action")
	ID := r.URL.Query().Get("id")
	commentID, _ := strconv.Atoi(ID)
	fmt.Print("comment like status of comment:", commentID)
	fmt.Println(action)
	switch action {
	case "like":
		fmt.Println("we liking a comment inna it")
		err := m.SaveCommentLike(commentID, user.ID)
		if err != nil {
			fmt.Println("error liking comment:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "removeLike":
		fmt.Println("we removing a comment like inna it")
		err := m.RemoveCommentLike(commentID, user.ID)
		if err != nil {
			fmt.Println("error removing comment like:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "dislike":
		fmt.Println("we disliking a comment inna it")
		err := m.DislikeComment(commentID, user.ID)
		if err != nil {
			fmt.Println("error disliking comment", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "removeDislike":
		fmt.Println("we removing a dislike from a comment inna it")
		err := m.RemoveCommentDislike(commentID, user.ID)
		if err != nil {
			fmt.Println("error removing comment dislike:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
