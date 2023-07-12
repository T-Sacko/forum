package routes

import (
	"net/http"
	c "forum/controllers"
)


func SetUpRoutes(mux *http.ServeMux) {
	
	
	fs := http.FileServer(http.Dir("static"))
	staticPrefix := "/static/"
	mux.Handle(staticPrefix, http.StripPrefix(staticPrefix, fs))  
	
	
	mux.HandleFunc("/", c.Index)
	mux.HandleFunc("/sign-in", c.UsersHandler)
	mux.HandleFunc("/api/check-username", c.UsernameCheck)
	mux.HandleFunc("/api/check-email", c.EmailCheck)
	mux.HandleFunc("/api/create-post", c.CheckSession)
	mux.HandleFunc("/create-post",c.CreatePost)
	mux.HandleFunc("/login",c.Login)
	mux.HandleFunc("/sign-up",c.SignUp)
	mux.HandleFunc("/like-post",c.LikePost)
	
}
