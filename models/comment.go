package models

import (
	"fmt"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

func (c *Comment) Save() (int64, error) {
	query := `insert into comments (content, post_id, user_id, created_at, updated_at) VALUES (?,?,?,?,?)`
	result, err := db.Exec(query, c.Content, c.PostID, c.UserID, c.CreatedAt, c.UpdatedAt)

	if err != nil{
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("the comment id is: ", ID)
	return ID, nil
}

// func GetCommentsByPostID(id int) ([]*Comment, error) {
//     // query the database for all comments on a given post
// }

// func GetCommentByID(id int) (*Comment, error) {
//     // query the database for a comment with the given ID
// }

// func GetCommentsByUserID(id int) ([]*Comment, error) {
//     // query the database for all comments made by a given user
// }
