package main

import (
	// "encoding/json"
	// "fmt"
	"forum/models"
	"forum/routes"
	"net/http"
	// _ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	models.InitDB()

	routes.SetUpRoutes()

	err := http.ListenAndServe(":5505", nil)
	if err != nil {
		panic(err)
	}
}
