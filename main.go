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

func main() {
	err := models.InitDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	fmt.Println("yoyo")
	mux := http.NewServeMux()
	r.SetUpRoutes(mux)

	fmt.Println("Serving on Port -> http://localhost:8888")
	if err := http.ListenAndServe("0.0.0.0:8888", mux); err != nil {
		log.Fatalf("Failure on Listening and Serving: %v", err)
	}
}
