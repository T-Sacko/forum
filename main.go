package main

import (
	"fmt"
	"forum/models"
	r "forum/routes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return os.ErrInvalid
	}

	return cmd.Start()
}

func main() {
	models.InitDB()

	defer models.CloseDB()

	// models.DeleteLikesTable()

	mux := http.NewServeMux()

	r.SetUpRoutes(mux)

	openBrowser("http://0.0.0.0:8888")
	fmt.Println("Serving on Port ->:8888")
	if err := http.ListenAndServe("0.0.0.0:8888", mux); err != nil {

		log.Fatalf("Failure on Listening and Serving: %v", err)
	}

}
