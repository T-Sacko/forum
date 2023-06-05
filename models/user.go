package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// func CreateUser(email, username, password string) {
// 	db.Exec("INSERT INTO users (email, username, password) VALUES (?,?,?)", email, username, password)
// }

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	SessionId string `json:"sessionid"`
}

func (u *User) Save() error {

	_, err := db.Exec("INSERT INTO users (email, username, password) VALUES (?,?,?)", u.Email, u.Username, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u User) IsUsernameAvailable() (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", u.Username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (u User) IsEmailAvailable() (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", u.Email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func SetSessionId(sessionId string) error {
	_, err := db.Exec("INSERT INTO users (sessionId) values(?)", sessionId)
	if err != nil {
		return err
	}
	return nil
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

// func GetUserByID(id int) (*User, error) {
//     // query the database for a user with the given ID
// }
