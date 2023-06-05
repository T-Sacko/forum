package models

import (
	"database/sql"
	"fmt"

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
	HashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, errs := db.Exec("INSERT INTO users (email, username, password) VALUES (?,?,?)", u.Email, u.Username, HashedPass)
	if errs != nil {
		return err
	}
	return nil
}

func CheckUserCredentials(email, password string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, username, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("sql erere")
			return user, fmt.Errorf("invalid username or password")
		}
		return user, err
	}

	if !ComparePasswords(user.Password, password) {
		fmt.Println("sql erere")

		return user, fmt.Errorf("invalid username or password")
	}

	return user, nil
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

func (u User) SetSessionId(sessionId string) error {

	available, _ := IsUsernameAvailable(u.Username)

	if available {

		_, err := db.Exec("INSERT INTO users (sessionId) values(?)", sessionId)
		if err != nil {
			return err
		}
		return nil
	}

	_, err := db.Exec("UPDATE users SET sessionId WHERE id = ?", u.ID)

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

func userExists(username string) bool {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// func GetUserByID(id int) (*User, error) {
//     // query the database for a user with the given ID
// }
