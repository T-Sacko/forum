package models

import (
	"fmt"
	"log"
	"time"
)

type Comment struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Content        string    `json:"content"`
	PostID         int       `json:"post_id"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Likes          int       `json:"likes"`
	Dislikes       int       `json:"dislikes"`
	UserLikeStatus int       `json:"user_like_status"`
}

func (c *Comment) Save() (int64, error) {
	query := `insert into comments (content, post_id, user_id) VALUES (?,?,?)`
	result, err := db.Exec(query, c.Content, c.PostID, c.UserID)

	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("the comment id is: ", ID)
	return ID, nil
}

func GetComments(userID, postID int) ([]Comment, error) {
	var comments []Comment

	query := `
	SELECT 
    comments.id, 
    users.username, 
    comments.content, 
    comments.post_id, 
    comments.user_id, 
    comments.created_at, 
    comments.updated_at,
    COALESCE(SUM(CASE WHEN comment_likes.value = 1 THEN 1 ELSE 0 END), 0) AS likes,
    COALESCE(SUM(CASE WHEN comment_likes.value = -1 THEN 1 ELSE 0 END), 0) AS dislikes,
    COALESCE(
        (SELECT value FROM comment_likes 
         WHERE comment_likes.comment_id = comments.id AND comment_likes.userId = ?), 
    0) AS user_like_status
FROM 
    comments
INNER JOIN 
    users ON comments.user_id = users.id
LEFT JOIN 
    comment_likes ON comment_likes.comment_id = comments.id
WHERE 
    comments.post_id = ?
GROUP BY 
    comments.id
ORDER BY 
    comments.id DESC	
    `

	rows, err := db.Query(query, userID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := Comment{}
		if err := rows.Scan(&comment.ID, &comment.Username, &comment.Content, &comment.PostID, &comment.UserID, &comment.CreatedAt, &comment.UpdatedAt, &comment.Likes, &comment.Dislikes, &comment.UserLikeStatus); err != nil {
			log.Println(err)
			continue
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(comments,"dddddddddddddddddddddddddddddddddddddddddddddddddddffffffffffffffffffffffffffffff")
	return comments, nil
}

// func GetCommentByID(id int) (*Comment, error) {
//     // query the database for a comment with the given ID
// }

// func GetCommentsByUserID(id int) ([]*Comment, error) {
//     // query the database for all comments made by a given user
// }
