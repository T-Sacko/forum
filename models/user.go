package models

// func CreateUser(email, username, password string) {
// 	db.Exec("INSERT INTO users (email, username, password) VALUES (?,?,?)", email, username, password)
// }


type User struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Username  string    `json:"username"`
    Password  string    `json:"-"`

}

func (u *User) Save() error {

	_, err := db.Exec("INSERT INTO users (email, username, password) VALUES (?,?,?)", u.Email, u.Username, u.Password)
	if err != nil {
	return err
	}
	return nil
}

// func GetUserByID(id int) (*User, error) {
//     // query the database for a user with the given ID
// }
