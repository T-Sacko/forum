package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	posts, err := m.GetPostsFromDB()
	if err != nil {
		// Handle the error (e.g., show an error page)
		fmt.Println("error with getposts")
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}

	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
	}

	data := struct {
		Posts []m.Post
	}{
		Posts: posts,
	}

	errs := Tpl.ExecuteTemplate(w, "home.html", data)
	fmt.Println("no sirr")
	if errs != nil {
		fmt.Println("no sir", errs)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func GetPostLikes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("a post ting working still")
	fmt.Println(r.Method, "the ,ethod is <<")
	userId, err := m.GetUserByCookie(r)
	if err != nil {
		fmt.Println("no cookie tring to get user liked posts",err)
		return
	}
	fmt.Println("getting user post likes for user ", userId, "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	likesData, err := m.GetLikedPosts(userId)
	if err!= nil{
		fmt.Println("error with suttin")
	}
	fmt.Println(likesData, "likes data is here u know")
	err1 := json.NewEncoder(w).Encode(likesData)
	if err1 != nil {
		fmt.Println("cant encode suttin")
	}
}
