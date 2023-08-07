package models

// import "time"

type Comment struct {
	ID int `json:"id"`
	Username string `json:"username"`
	UserID   int    `json:"userId"`
	PostID   string `json:"postId"`
	Comment  string `json:"comment"`
}

func (comment Comment) SaveComment() error {
	stmt, err := db.Prepare("INSERT INTO comments (content, postId, userId) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Execute the prepared statement with the provided values and current timestamp
	_, err = stmt.Exec(comment.Comment, comment.PostID, comment.UserID)
	if err != nil {
		return err
	}

	return nil
}

// func getComments(ID int) (string, error) {
// 	stmt, err := db.Prepare("SELECT content FROM comments WHERE id = ?")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer stmt.Close()

// 	var comments string
// 	// Assuming "idValue" is the ID of the comment you want to retrieve
// 	err = stmt.QueryRow(ID).Scan(&comments)
// 	if err != nil {
// 		return "", err
// 	}
// 	return comments, nil
// }

// type Comment struct {
//     ID          int       `json:"id"`
//     Content     string    `json:"content"`
//     PostID      int       `json:"post_id"`
//     UserID      int       `json:"user_id"`
//     CreatedAt   time.Time `json:"created_at"`
//     UpdatedAt   time.Time `json:"updated_at"`
// }

// func (c *Comment) Save() error {
//     // save the comment to the database
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
