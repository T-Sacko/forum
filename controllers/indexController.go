package controllers

import (
	"fmt"
	m "forum/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 doesnt exist", http.StatusNotFound)
		return
	}
	var internal = "INTERNAL SERVER ERROR:"
	var user m.User
	filter := "home"
	user1, _ := m.GetUserByCookie(r)
	if user1 != nil {
		user = *user1
	}
	var posts []m.Post
	var err error
	category := r.URL.Query().Get("category")
	inErr := fmt.Sprintf("%s Failed to retrieve posts", internal)
	if category == "liked-posts" {
		filter = category
		posts, err = m.FilterByLiked(user.ID)
		if err != nil {
			// Handle the error (e.g., show an error page)
			fmt.Println("error with getposts")
			http.Error(w, inErr, http.StatusInternalServerError)
			return
		}
	} else if category == "my-posts" {
		filter = category
		posts, err = m.FilterByUserPosts(user.ID)
		if err != nil {
			// Handle the error (e.g., show an error page)
			fmt.Println("error with getposts")
			http.Error(w, inErr, http.StatusInternalServerError)
			return
		}
	} else if category != "" {
		fmt.Println("category is:", category)
		filter = category
		posts, err = m.FilterByCategory(user.ID, category)
		fmt.Println("lenght is :", len(posts), posts)
		if err != nil {
			// Handle the error (e.g., show an error page)
			fmt.Println("error with getpostsFilter")
			http.Error(w, inErr, http.StatusInternalServerError)
			return
		}
	} else {
		posts, err = m.GetPostsFromDB(user.ID)
		if err != nil {
			// Handle the error (e.g., show an error page)
			fmt.Println("error with getposts")
			http.Error(w, inErr, http.StatusInternalServerError)
			return
		}
	}

	// var category string

	// for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
	// 	posts[i], posts[j] = posts[j], posts[i]
	// }
	categories := []string{"biology", "etymology"}
	data := struct {
		Posts      []m.Post
		Categories []string
		User       m.User
		Filter     string
	}{
		Posts:      posts,
		Categories: categories,
		User:       user,
		Filter:     filter,
	}

	errs := Tpl.ExecuteTemplate(w, "home.html", data)

	if errs != nil {
		fmt.Println("no sir", errs)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

// func GetPostLikes(w http.ResponseWriter, r *http.Request) {

// 	user, err := m.GetUserByCookie(r)
// 	if err != nil {
// 		fmt.Println("no liked posts", err)
// 		return
// 	}
// 	likesData, err := m.GetLikedPosts(user.ID)
// 	if err != nil {
// 		fmt.Println("error with suttin")
// 	}
// 	err1 := json.NewEncoder(w).Encode(likesData)
// 	if err1 != nil {
// 		fmt.Println("cant encode suttin")
// 	}
// }
