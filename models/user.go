package models

import (
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
	Post  Content
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

	userInfo, err :=  user.GetUserByID()
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
		Post: content,
		Likes: likes,
		Dislikes: dislikes,
		Comments: comments,
	}
	
	return user, err
}





// func (u *User) Save() error {
//     // save the user to the database
// }

// func GetUserByID(id int) (*User, error) {
//     // query the database for a user with the given ID
// }
