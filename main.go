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

func Start() {
	main()
}

func main() {
	models.InitDB()
	defer models.CloseDB()
	LocalHost := ":8888"
	url := "https://localhost" +LocalHost
	exec.Command("open", url).Start()
	mux := http.NewServeMux()
	r.SetUpRoutes(mux)
	err := http.ListenAndServe(LocalHost, mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Serving on Port ->:8888")
}
