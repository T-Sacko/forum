package main

import (
	"fmt"
	c "forum/controllers"
	"forum/models"
	r "forum/routes"
	"html/template"
	"log"
	"net/http"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	c.Tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func openBrowser(url string) error {
	return exec.Command("xdg-open", url).Start()
}

func main() {
	models.InitDB()

	defer models.CloseDB()

	mux := http.NewServeMux()
	r.SetUpRoutes(mux)
	
	openBrowser("http://localhost:8888")
	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatalf("Failure on Listening and Serving: %v", err)
	}


	fmt.Println("Serving on Port ->:8888")
}
