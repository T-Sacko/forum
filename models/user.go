package models

func CreateUser(email, username, password string) {
	db.Exec("INSERT INTO users (email, username, password) VALUES (?,?,?)", email, username, password)
}

// import "time"

// type User struct {
//     ID        int       `json:"id"`
//     Email     string    `json:"email"`
//     Username  string    `json:"username"`
//     Password  string    `json:"-"`
//     CreatedAt time.Time `json:"created_at"`
//     UpdatedAt time.Time `json:"updated_at"`
// }

// func (u *User) Save() error {
//     // save the user to the database
// }

// func GetUserByID(id int) (*User, error) {
//     // query the database for a user with the given ID
// }
