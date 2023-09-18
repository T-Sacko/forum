package main

import (
	"fmt"
	"forum/models"
	r "forum/routes"
	"log"
	"net/http"

	//"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

// func init() {
// 	c.Tpl = template.Must(template.ParseGlob("templates/*.html"))
// }

// func openBrowser(url string) error {
// 	return exec.Command("xdg-open", url).Start()
// }

func main() {
	models.InitDB()

	//defer models.CloseDB()fmt
	fmt.Println("yoyo")

	mux := http.NewServeMux()
	r.SetUpRoutes(mux)
	// posts, _ := models.GetPostsFromDB()
	// fmt.Println(posts)

	// openBrowser("http://0.0.0.0:8888")
	fmt.Println("Serving on Port -> http://localhost:8888")
	if err := http.ListenAndServe(":8888", mux); err != nil {

		log.Fatalf("Failure on Listening and Serving: %v", err)
	}

}
