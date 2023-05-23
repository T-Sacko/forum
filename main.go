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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func index(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		HashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		err = models.InsertDB(username, email, string(HashedPass))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "sign-in.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "create_post.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Front-End/home.html")
	if err != nil {
		panic(err)
	}

func main() {
	models.InitDB()
	defer models.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/static/", staticHandler) // Add this line to handle the /static/ route
	http.Handle("/static/", http.StripPrefix("/static/", mux))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign-in", UsersHandler)
	mux.HandleFunc("/posts", PostsHandler)
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		panic(err)
	}
}
