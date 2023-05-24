package main

import (
	// "encoding/json"
	// "fmt"
	"forum/models"
	"forum/routes"
	"net/http"
	"os/exec"
	// _ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	models.InitDB()
	exec.Command("open",)
	routes.SetUpRoutes()

	err := http.ListenAndServe(":5505", nil)
	if err != nil {
		panic(err)
	}
}
