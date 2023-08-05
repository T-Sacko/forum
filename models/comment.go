package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Likes      int		`json:"likes"`
	Dislikes   int		`json:"dislikes"`
}

// func (c *Comment) Save() error {
//	query := `insert into comments (content, post_id, user_id, created_at, updated_at) VALUES (?,?,?,?,?)`
//   _, err :=  db.exec(query, c.content, c.post_id, user_id, created_at, updated_at)
//   return err
// }

// func GetCommentByID(id int) (*Comment, error) {
//     // query the database for a comment with the given ID
// }

// func GetCommentsByPostID(id int) ([]*Comment, error) {
//     // query the database for all comments on a given post
// }

// func GetCommentsByUserID(id int) ([]*Comment, error) {
//     // query the database for all comments made by a given user
// }
