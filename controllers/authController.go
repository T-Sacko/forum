package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	m "forum/models"
	"io"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var err error

// var Tpl = template.Must(template.ParseGlob("templates/*.html"))
var state = uuid.Must(uuid.NewV4()).String()

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8888/google",
	ClientID:     "128618599322-8r8rlh7nqj32kgdt8napmiji1q8vei31.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

// func HandleGoogleAuth(w http.ResponseWriter, r *http.Request) {
// 	url := googleOauthConfig.AuthCodeURL(oauthStateString)
// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// }

func GoogleAuth(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "select_account"))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println(state)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("callbck received", r.URL.Query().Get("state"))
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// Step 2: Use the code to exchange for an access token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)

	// Step 2: Use this client to make a request to the userinfo endpoint
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Failed to read user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var userInfo map[string]interface{}
	err = json.Unmarshal(data, &userInfo)
	if err != nil {
		http.Error(w, "Error unmarshalling user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// email, ok := userInfo["email"].(string)

	// userInfo now contains the user's information.
	// For demonstration, we're just printing it out:
	fmt.Printf("User Infos: %+v\n", userInfo["email"])
}
func Session(w http.ResponseWriter, r *http.Request) {
	var seshData struct {
		Status bool `json:"status"`
	}
	_, err := m.GetUserByCookie(r)
	if err != nil {
		json.NewEncoder(w).Encode(seshData)
		return
	}
	seshData.Status = true
	json.NewEncoder(w).Encode(seshData)

}

// func cookie(r *http.Request) (string, error) {
// 	cookie, err := r.Cookie("session")
// 	if err != nil {
// 		return "", err
// 	}
// 	seshID := cookie.Value
// 	return seshID, nil

// }

// ////////////////////////////////////////////////
func SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("log out req received")
	seshID := r.URL.Query().Get("seshID")
	err = m.DeleteCookie(seshID)
	if err != nil {
		fmt.Println("cant log out:", err)
		w.WriteHeader(http.StatusUnauthorized)
	}

}

// ////////////////////////////////////////////////////
func Login(w http.ResponseWriter, r *http.Request) {

	var user m.User = m.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	fmt.Println("login req received", user.Email, user.Password)

	isUser, _ := m.Check4User(user.Email, user.Password)
	if !isUser {
		fmt.Println(isUser, user.Password, "not nil")

		http.Error(w, "unauth", http.StatusUnauthorized)
		return
	}
	// Set the session ID and create a cookie
	SessionId, err := uuid.NewV1()
	if err != nil {
		http.Error(w, "Couldn't create sesh id", http.StatusInternalServerError)
		return
	}
	m.SetSessionId(user.Email, SessionId.String())
	cookie := &http.Cookie{
		Name:  "session",
		Value: SessionId.String(),
		// Session cookie (valid until browser is closed)
	}

	http.SetCookie(w, cookie)
	fmt.Println("yessssss")
	http.Redirect(w, r, "/", http.StatusFound)

}

func SignUp(w http.ResponseWriter, r *http.Request) {

	fmt.Println("signup req received")
	SessionId, err := uuid.NewV1()
	if err != nil {
		http.Error(w, "ERROR 500", http.StatusInternalServerError)
		return
	}
	user := getUser(r)
	user.SessionId = SessionId.String()
	err = user.Register()
	if err != nil {
		http.Error(w, "CANT SAVE USER", http.StatusBadRequest)
		http.Redirect(w, r, "", http.StatusSeeOther)
		return
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: user.SessionId,
	}
	fmt.Println("man", cookie.Value)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)

}

func getUser(r *http.Request) *m.User {
	return &m.User{Email: r.FormValue("email"), Username: r.FormValue("username"), Password: r.FormValue("password")}
}

// func UsersHandler(w http.ResponseWriter, r *http.Request) {
// 	err = Tpl.ExecuteTemplate(w, "sign-in.html", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// ----------------------------------------------------ajax

func UsernameCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	available, err := m.IsUsernameAvailable(username)
	fmt.Println("the username", username, "is available", available)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response := m.UserCheckResponse{
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
func EmailCheck(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	available, err := m.IsEmailAvailable(email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response := m.UserCheckResponse{
		Available: available,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("can't encode email check response into json", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
