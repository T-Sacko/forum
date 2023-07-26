package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
	"strconv"
)

func HandlePostLikes(w http.ResponseWriter, r *http.Request) {

	sessionId, _ := checkCookie(w, r)
	userId, session, err1 := m.SessionIsActive(sessionId)
	if err1 != nil {
		fmt.Println("post.controller.go Error func HandlePostAction: ", err1)
		return
	}

	if session {
		postActionReqData := m.PostActionReq{}
		json.NewDecoder(r.Body).Decode(&postActionReqData)
		fmt.Println("likeActionData una", postActionReqData)
		postId, _ := strconv.Atoi(postActionReqData.PostId)
		action := postActionReqData.Action

		switch action {
		case "like":
			m.SaveLike(postId, userId)
		case "unlike":
			m.RemoveLike(postId, userId)
		case "dislike":
			m.SaveDislike(postId, userId)
		case "removeDislike":
			m.RemoveDislike(postId, userId)
		}
	}
}
