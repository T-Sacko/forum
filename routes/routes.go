package routes

import (
	"forum/controllers"
	"net/http"
)

func SetUpRoutes(){

http.HandleFunc("/",controllers.HomePage)

http.HandleFunc("/register",controllers.RegisterHandler)

http.HandleFunc("/login",controllers.LoginHandler)

}