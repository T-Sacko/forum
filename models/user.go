package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserCheckResponse struct {
	ID        int  `json:"id"`
	Available bool `json:"available"`
	UserInfo  User
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	SessionId string `json:"sessionid"`
	Post     string
	Likes    int
	Dislikes int
	Comments string
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
	likes, err := getLikes(user.ID)
	if err != nil {
		return User{}, err
	}

	dislikes, err := getDislikes(user.ID)
	if err != nil {
		return User{}, err
	}

	comments, err := getComments(user.ID)
	if err != nil {
		return User{}, err
	}

	user = User{
		
		Likes:    likes,
		Dislikes: dislikes,
		Comments: comments,
	}

	return user, err
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
		if err == sql.ErrNoRows {
			return false, err
		} else {
			return false, err
		}
	}
	return ComparePasswords(hashedPassword, password), nil
}
