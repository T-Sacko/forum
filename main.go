package main

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var userDetails = `
<html>
<head>
  <title>User Details</title>
</head>
<body>
  <h1>User Details</h1>

  <h2>User Information</h2>
  <div>
    <strong>Username:</strong> <span id="username"></span><br>
    <strong>Email:</strong> <span id="email"></span>
  </div>

  <h2>Posts</h2>
  <ul id="posts"></ul>

  <script>
    // Assuming you have a JavaScript object containing user data and posts
    var userData = {
      username: "JohnDoe",
      email: "johndoe@example.com"
    };

    var userPosts = [
      { id: 1, title: "Post 1", content: "Lorem ipsum dolor sit amet." },
      { id: 2, title: "Post 2", content: "Pellentesque habitant morbi tristique senectus." },
      { id: 3, title: "Post 3", content: "Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae." }
    ];

    // Update user information
    document.getElementById("username").textContent = userData.username;
    document.getElementById("email").textContent = userData.email;

    // Render user posts
    var postsElement = document.getElementById("posts");
    userPosts.forEach(function(post) {
      var li = document.createElement("li");
      var title = document.createElement("h3");
      var content = document.createElement("p");

      title.textContent = post.title;
      content.textContent = post.content;

      li.appendChild(title);
      li.appendChild(content);
      postsElement.appendChild(li);
    });
  </script>
</body>
</html>
`

var err error

type Post struct {
	ID       int
	Username string
	Email    string
	Content  string
	Likes    int
}

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
	if r.URL.Path == "/account" {
		tpl = template.Must(template.New("userDetails").Parse(userDetails))
		//Open DB for the User's information
		tpl.Execute(w, nil)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "sign-in.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println(username, email, password)
		HashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		err = models.InsertDB(username, email, string(HashedPass))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "create_post.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func UsernameCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	available, err := models.IsUsernameAvailable(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := models.UsernameCheckResponse{
		Available: available,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	models.InitDB()

	defer models.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/static/", staticHandler) // Add this line to handle the /static/ route
	http.Handle("/static/", http.StripPrefix("/static/", mux))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign-in", UsersHandler)
	mux.HandleFunc("/api/check-username", UsernameCheck)
	mux.HandleFunc("/posts", PostsHandler)
	
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Serving on Port ->:8888")
}
