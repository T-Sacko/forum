package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	//"os/exec"
)

var tpl *template.Template

const LoggedIn = `
<style>
.register {
	background-color: white;
 border: cornflowerblue;
	color: rgb(0, 183, 255);
 padding: 12px 16px;
	font-size: 16px;
	cursor: pointer;
}
</style>
<body style="background-color: cadetblue;">
   <u><p><h1 align = "center">Welcome to Forum, {{.Name}}</h1></p>
   <div align = "center">
   <form action="/Forum" method= "POST">
    <p><textarea name="Post" cols="50" rows="20" placeholder="Write something here..."></textarea></p>
    <button type="submit" class = "register">Post!</button>

`

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

type LogIn struct {
	Name string
	Post string
	Time Date
}

var b LogIn

type register int

func (rg register) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		tpl = template.Must(template.ParseFiles("register.html"))
		tpl.Execute(w, nil)
	case "/Forum":
		if req.FormValue("Username") == "" || req.FormValue("Password") == "" && req.FormValue("Email") == "" {
			/*if req.FormValue("Post") != "" {
				b.Name = req.FormValue("Username")
				b.Post = req.FormValue("Post")
				b.Time.Year, b.Time.Month, b.Time.Day = time.Now().Date()
				tpl = template.Must(template.New("UrPosts").Parse(UrPosts))
				tpl.Execute(w, b)
				return
			}*/
			tpl = template.Must(template.ParseFiles("Index.html"))
			tpl.Execute(w, nil)
		} else {
				b.Name = req.FormValue("Username")
				tpl = template.Must(template.New("homePage").Parse(LoggedIn))
				tpl.Execute(w, b)
		}
	}
}

func main() {
	var rg register
	err := http.ListenAndServe(":6789", rg)
	if err != nil {
		log.Fatal(err.Error())
	}
}

//var mainPage = "6789/Forum"
/*
var mux = http.NewServeMux()


func main() {
	mux.HandleFunc("/Forum", Index)
	mux.HandleFunc("/", register)
	//	http.HandleFunc("/homePage", HP)
	err := http.ListenAndServe(":6789", mux)
	if err != nil {
		fmt.Println("yeah")
		log.Fatal(err.Error())
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
		tpl = template.Must(template.ParseFiles("Index.html"))
		tpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("register.html"))
	tpl.Execute(w, nil)
}*/
