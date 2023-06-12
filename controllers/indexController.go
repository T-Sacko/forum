package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("DONE")
		switch r.URL.Path {
		case "/log-in":
			auth, check := SigningIn(w, r)
			if auth == "Authenticated" {
				fmt.Println(auth)
				err = Tpl.ExecuteTemplate(w, "home.html", check.UserInfo)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				http.Error(w, auth, http.StatusNoContent)
			}
		default:
			err = getUser(r).Register()
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(w, r, "/", http.StatusAccepted)
		}
	}
	err = Tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
