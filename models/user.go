package models

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserCheckResponse struct {
	ID        int  `json:"id"`
	Available bool `json:"available"`
	UserInfo  User
}

type User struct {
	Status    bool
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	SessionId string `json:"sessionid"`
	Post      string
	Likes     int
	Dislikes  int
	Comments  string
}

func DeleteCookie(seshID string) error {
	_, err := db.Exec("UPDATE users SET sessionID = NULL WHERE sessionId = ?", seshID)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByCookie(r *http.Request) (*User, error) {
	cookie, errs := r.Cookie("session")
	if errs != nil {
		fmt.Println("theres no cookie")
		return nil, errs
	}
	sessionId := cookie.Value
	var user User
	err := db.QueryRow("SELECT id, username, email FROM users WHERE sessionId = ?", sessionId).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		// Handle the database query error accordingly
		return nil, err
	}
	user.Status = true
	return &user, nil
}

func (newUser User) Register() error {
	Password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("its whre")

		return err
	}
	err = InsertDB(newUser.Username, newUser.Email, string(Password), newUser.SessionId)
	if err != nil {
		fmt.Println("cant sign up user alredy exists")
		return err
	}
	return nil
}

func (user User) LogIn() (UserCheckResponse, error) {
	checked, err := Check4User(user.Email, user.Password)
	if err != nil {
		return UserCheckResponse{Available: checked}, err
	}

	userInfo, err := user.GetUserByID()
	if err != nil {
		return UserCheckResponse{}, err
	}
	return UserCheckResponse{Available: checked, UserInfo: userInfo}, nil
}

func (user User) GetUserByID() (User, error) {
	// ID, err := GetID(user.Email)
	// if err != nil {
	// 	return User{}, err
	// }
	// user.ID = ID
	// content, err := getUserContent(user.ID)
	// if err != nil {
	// 	return User{}, err
	//
	// likes, err := getLikes(user.ID)
	// if err != nil {
	// 	return User{}, err
	// }

	// dislikes, err := getDislikes(user.ID)
	// if err != nil {
	// 	return User{}, err
	// }

	// comments, err := getComments(user.ID)
	// if err != nil {
	// 	return User{}, err
	// }

	// user = User{

	// 	Likes:    likes,
	// 	Dislikes: dislikes,
	// 	Comments: comments,
	// }

	return user, nil
}

func SetSessionId(email, sessionId string) error {
	_, err := db.Exec("UPDATE users SET sessionId = ? WHERE email = ?", sessionId, email)
	if err != nil {
		fmt.Println("failed to update session ID:", err)
		return err
	}
	return nil
}

func IsUsernameAvailable(username string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func IsEmailAvailable(email string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func ComparePasswords(hashedPassword, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	return err == nil
}

func Check4User(email, password string) (bool, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword)
	if err != nil {

		return false, err

	}
	return ComparePasswords(hashedPassword, password), nil
}
