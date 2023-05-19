package routes

import(
	"net/http"
)

func RegisterHandler(w *http.ResponseWriter r *http.Request) {
	err := r.ParseForm()
	if err != nil    {
		http.Error(w, "Failed to parseform")
	}

	email:= r.FormValue("email-address")
	username:=r.FormValue("username")
	password:=r.FormValue("password")
	

}