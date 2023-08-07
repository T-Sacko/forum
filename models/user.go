package models

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserCheckResponse struct {
	Available bool `json:"available"`
	UserInfo  *User
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	SessionId string `json:"sessionid"`
	Post      []Post
}

func getUsernameByID(Id int) string {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = ?", Id).Scan(&username)
	if err != nil {
		log.Fatal(err)
	}
	return username
}

func (newUser User) Register() error {
	Password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = InsertDB(newUser.Username, newUser.Email, string(Password), newUser.SessionId)
	if err != nil {
		return err
	}
	return nil
}

func SetSessionId(email, sessionId string) error {
	_, err := db.Exec("UPDATE users SET sessionId = ? WHERE email = ?", sessionId, email)
	if err != nil {
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
