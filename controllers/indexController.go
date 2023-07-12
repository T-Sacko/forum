package controllers

import (
	"fmt"
	m "forum/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	userId, err := m.GetUserByCookie(r)
	if err != nil {
		fmt.Println("nah g")
		
	}
	likesData, _ := m.GetLikedPosts(userId)
	fmt.Println(likesData, "likes data is here u know")
	
	posts, err := m.GetPostsFromDB()
	if err != nil {
		// Handle the error (e.g., show an error page)
		fmt.Println("error with getposts")
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}
	
	data := struct {
		Posts []m.Post
		LikesData []m.LikeData
	}{
		Posts: posts,
		LikesData: likesData,
	}

	errs := Tpl.ExecuteTemplate(w, "home.html", data)
	fmt.Println("no sirr")
	if errs != nil {
		fmt.Println("no sir", errs)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
