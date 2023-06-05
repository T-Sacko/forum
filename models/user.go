package models

import (
	"database/sql"
	"fmt"
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

	_, errs := db.Exec("INSERT INTO users (email, username, password) VALUES (?,?,?)", u.Email, u.Username, string(u.Password), )
	if errs != nil {
		return errs
	}
	return nil
}

func CheckUserCredentials(email, password string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, username, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("sql can qury to check creds")
			return user, fmt.Errorf("invalid username or password")
		}
		return user, err
	}

	if !ComparePasswords(user.Password, password) {
		fmt.Println("sql1 erere")

		return user, fmt.Errorf("invalid username or password")
	}

	return user, nil
}

func ComparePasswords(hashedPassword, userPassword string) bool {

	db.QueryRow()
	return hashedPassword == userPassword
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
		fmt.Println("cant check if username available")
		return false, err
	}
	return count == 0, nil
}

func (u User) SetSessionId(sessionId string) error {

	_, err := db.Exec("INSERT INTO users (sessionId) values(?)", sessionId)
	if err != nil {
		fmt.Println("cant insert sesh id")
		return err
	}
	return nil

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
