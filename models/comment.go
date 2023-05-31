package models

// import "time"

func getComments(ID int) (string, error) {
	stmt, err := db.Prepare("SELECT content FROM comments WHERE id = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var comments string
	// Assuming "idValue" is the ID of the comment you want to retrieve
	err = stmt.QueryRow(ID).Scan(&comments)
	if err != nil {
		return "", err
	}
	return comments, nil
}

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
