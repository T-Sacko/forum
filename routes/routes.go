package routes

import (
	"forum/controllers"
	"net/http"
)

func main(){

http.HandleFunc("/register",controllers.RegisterHandler)

}