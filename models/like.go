package models

import "fmt"

type LikeData struct {
	PostId int `json:"postId"`
	Value  int `json:"value"`
}

type PostActionReq struct {
	PostId string `json:"postId"`
	Action string `json:"action"`
}

func GetLikedPosts(userId int) ([]LikeData, error) {
	query := `
		SELECT postId, value
		FROM likes
		WHERE userId = ?
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		fmt.Println("error like.go", err)
		return nil, err
	}

	likesData := make([]LikeData, 0)

	for rows.Next() {

		var PostId, value int
		err := rows.Scan(&PostId, &value)
		if err != nil {
			return nil, err
		}

		LikeData := LikeData{
			PostId: PostId,
			Value:  value,
		}

		likesData = append(likesData, LikeData)

	}

	return likesData, nil

}

//------------------------------------------post likes

func SaveLike(postId, userId int) {
	RemoveDislike(postId,userId)
	_, err := db.Exec("INSERT INTO likes (postId, userId, value) Values (?, ?, ?)", postId, userId, 1)
	if err != nil {
		fmt.Println("like.go error inserting like: ", err)
	}
}

func RemoveLike(postId, userId int) {
	_, err := db.Exec("DELETE FROM likes WHERE postId = ? AND userId = ? AND value = 1", postId, userId)
	if err != nil {
		fmt.Println("like.go error removing like: ", err)
	}
}

func SaveDislike(postId, userId int) {
	RemoveLike(postId,userId)
	_, err := db.Exec("INSERT INTO likes (postId, userId, value) VALUES (?, ?, ?)", postId, userId, -1)
	if err != nil {
		fmt.Println("like.go error inserting dislike: ", err)
	}
}

func RemoveDislike(postId, userId int) {
	_, err := db.Exec("DELETE FROM likes WHERE postId = ? AND userId = ? AND value = -1", postId, userId)
	if err != nil {
		fmt.Println("like.go error removing dislike: ", err)
	}
}

//----------------------------------------------comment likes

func SaveCommentLike(commentID, userId int) error {
	RemoveDislike(commentID,userId)
	_, err := db.Exec("INSERT INTO comment_likes (comment_id, userId, value) Values (?, ?, ?)", commentID, userId, 1)
	if err != nil {
		fmt.Println("like.go error inserting like: ", err)
		return err
	}
	return nil
}

func RemoveCommentLike(commentID, userId int) error {
	_, err := db.Exec("DELETE FROM comment_likes WHERE comment_id = ? AND userId = ? AND value = 1", commentID, userId)
	if err != nil {
		fmt.Println("like.go error removing like: ", err)
		return err
	}
	return nil
}

func DislikeComment(commentID, userId int) error {
	RemoveLike(commentID,userId)
	_, err := db.Exec("INSERT INTO comment_likes (comment_id, userId, value) VALUES (?, ?, ?)", commentID, userId, -1)
	if err != nil {
		fmt.Println("like.go error inserting dislike: ", err)
		return err
	}
	return nil
}

func RemoveCommentDislike(commentID, userId int) error {
	_, err := db.Exec("DELETE FROM comment_likes WHERE comment_id = ? AND userId = ? AND value = -1", commentID, userId)
	if err != nil {
		fmt.Println("like.go error removing cDislike: ", err)
		return err
	}
	return nil
}
