package routes

import (
	"net/http"
	c "forum/controllers"
)


func SetUpRoutes(mux *http.ServeMux) {
	
	http.Handle("/static/", http.StripPrefix("/static/", mux))
	mux.HandleFunc("/", c.Index)
	mux.HandleFunc("/sign-in", c.UsersHandler)
	mux.HandleFunc("/api/check-username", c.UsernameCheck)
	mux.HandleFunc("/api/check-email", c.EmailCheck)
	mux.HandleFunc("/api/create-post", c.CheckSession)
	mux.HandleFunc("/create-post",c.CreatePost)
	mux.HandleFunc("/login",c.Login)
	mux.HandleFunc("/sign-up",c.SignUp)
	
}
