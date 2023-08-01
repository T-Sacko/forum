package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var user m.User
	user1, _ := m.GetUserByCookie(r)
	if user1!=nil{
		user = *user1
	}
	var posts []m.Post
	var err error
	category := r.URL.Query().Get("category")
	if category == "liked-posts" {
		posts, err = m.FilterByLiked(user.ID)
		if err != nil {
			// Handle the error (e.g., show an error page)
			fmt.Println("error with getposts")
			http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
			return
		}
	} else if category != "" {
		posts, err = m.FilterByCategory(category)
		if err != nil {
			// Handle the error (e.g., show an error page)
			fmt.Println("error with getposts")
			http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
			return
		}
	} else {

		posts, err = m.GetPostsFromDB()
		if err != nil {
			// Handle the error (e.g., show an error page)
			fmt.Println("error with getposts")
			http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
			return
		}
	}

	// var category string

	// for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
	// 	posts[i], posts[j] = posts[j], posts[i]
	// }
	categories := []string{"biology", "etymology", "sociology"}
	data := struct {
		Posts      []m.Post
		Categories []string
		User       m.User
	}{
		Posts:      posts,
		Categories: categories,
		User:       user,
	}

	errs := Tpl.ExecuteTemplate(w, "home.html", data)

	if errs != nil {
		fmt.Println("no sir", errs)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func GetPostLikes(w http.ResponseWriter, r *http.Request) {
	
	user, err := m.GetUserByCookie(r)
	if err != nil {
		fmt.Println("no cookie tring to get user liked posts", err)
		return
	}
	likesData, err := m.GetLikedPosts(user.ID)
	if err != nil {
		fmt.Println("error with suttin")
	}
	err1 := json.NewEncoder(w).Encode(likesData)
	if err1 != nil {
		fmt.Println("cant encode suttin")
	}
}
