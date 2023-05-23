package main

import (
	// "encoding/json"
	// "fmt"
	// "forum/models"
	"html/template"
	"log"
	"net/http"
	// _ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/bcrypt"
)

// var err error

// type Post struct {
// 	ID       int
// 	Username string
// 	Email    string
// 	Content  string
// 	Likes    int
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	tpl := template.Must(template.ParseGlob("templates/*.html"))
// 	err = tpl.ExecuteTemplate(w, "home.html", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if r.Method == "POST" {
// 		username := r.FormValue("username")
// 		email := r.FormValue("email")
// 		password := r.FormValue("password")
// 		HashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		err = models.InsertDB(username, email, string(HashedPass))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	fmt.Println("Data stored in the database")

// }

// func UsernameCheck(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		username := r.URL.Query().Get("username")
// 		available, err := models.IsUsernameAvailable(username)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		response := struct {
// 			Available bool `json:"available"`
// 		}{
// 			Available: available,
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 	}
// }

func main() {
	// models.InitDB()
	// defer models.CloseDB()
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":5505", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Front-End/home.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
