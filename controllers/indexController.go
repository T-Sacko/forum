package controllers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	Tpl.ExecuteTemplate(w, "home.html", nil)
}
