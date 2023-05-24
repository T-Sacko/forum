package main

import (
	"fmt"
	"forum/models"
	"html/template"
	"log"
	"net/http"
	c "forum/controllers"
	_ "github.com/mattn/go-sqlite3"
	r "forum/routes"
)


type Post struct {
	ID       int
	Username string
	Email    string
	Content  string
	Likes    int
}


func init() {
	c.Tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	models.InitDB()
	defer models.CloseDB()

	mux := http.NewServeMux()
	r.SetUpRoutes(mux)
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Serving on Port ->:8888")
}
