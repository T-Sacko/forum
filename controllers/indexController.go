package controllers

import (
	"fmt"
	m "forum/models"
	"html/template"
	"net/http"
)
var tol *template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	posts, err := m.GetPostsFromDB()
	if err != nil {
		// Handle the error (e.g., show an error page)
		fmt.Println("error with getposts")
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}

	data := struct {
		Posts []m.Post
	}{
		Posts: posts,
	}
 tol,_=template.ParseFiles("../templates/home.html")
	errs := tol.Execute(w, data)
	
	if errs != nil {
		fmt.Println("no sir", errs)
		// http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
