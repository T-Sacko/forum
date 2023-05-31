package routes

import (
	"net/http"
	c "forum/controllers"
)


func SetUpRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/static/", c.StaticHandler) // Add this line to handle the /static/ route
	http.Handle("/static/", http.StripPrefix("/static/", mux))
	mux.HandleFunc("/", c.Index)
	mux.HandleFunc("/sign-in", c.UsersHandler)
	mux.HandleFunc("/api/check-username", c.UsernameCheck)
	mux.HandleFunc("/api/check-email", c.EmailCheck)
	mux.HandleFunc("/posts", c.PostsHandler)
}
