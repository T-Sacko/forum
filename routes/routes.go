package routes

import (
	"net/http"
	c "forum/controllers"
)


func SetUpRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/static/", c.StaticHandler)
	http.Handle("/static/", http.StripPrefix("/static/", mux))
	mux.HandleFunc("/", c.Index)
	mux.HandleFunc("/api/check-username", c.UsernameCheck)
	mux.HandleFunc("/api/check-email", c.EmailCheck)
	mux.HandleFunc("/api/create-post", c.CheckSession)
}
