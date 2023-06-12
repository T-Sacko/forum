package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserCheckResponse struct {
	ID        int  `json:"id"`
	Available bool `json:"available"`
	UserInfo  User
}

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Post     Content
	Likes    int
	Dislikes int
	Comments string
}

// import "time"


func (newUser User) Register() error {
	HashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = InsertDB(newUser.Username, newUser.Email, string(HashedPass))
	if err != nil {
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
	ID, err := GetID(user.Email)
	if err != nil {
		return User{}, err
	}
	user.ID = ID
	content, err := getContent(user.ID)
	if err != nil {
		return User{}, err
	}
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
		Post:     content,
		Likes:    likes,
		Dislikes: dislikes,
		Comments: comments,
	}

	return user, err
}

func InsertDB(username, email, password string) error {
	// Prepare the SQL statement to insert a new post
	stmt, err := db.Prepare("INSERT INTO users (username, email, password, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Get the current timestamp
	currentTime := time.Now()

	// Execute the prepared statement with the provided values and current timestamp
	_, err = stmt.Exec(username, email, password, currentTime)
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


