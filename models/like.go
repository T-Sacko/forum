package models

import "fmt"

type LikeData struct {
	Likes    []Likes    `json:"likes"`
	Dislikes []Dislikes `json:"dislikes"`
}

type Likes struct {
	CommentId int `json:"commentId"`
	PostId    int `json:"postId"`
	Value     int `json:"value"`
}

type Dislikes struct {
	CommentId int `json:"commentId"`
	PostId    int `json:"postId"`
	Value     int `json:"value"`
}

type PostActionReq struct {
	PostId    string `json:"postId"`
	Action    string `json:"action"`
	CommentId string `json:"commentId"`
}

func GetLikedPosts(postId int) (LikeData, error) {
	var likesData LikeData

	query := `
		SELECT userId, comment_id, value
		FROM likes
		WHERE postId = ?
	`
	rows, err := db.Query(query, postId)
	if err != nil {
		fmt.Println("error like.go", err)
		return LikeData{}, err
	}
	likes := []Likes{}
	for rows.Next() {
		var PostId, CommentId, value int
		err := rows.Scan(&PostId, &CommentId, &value)
		if err != nil {
			return LikeData{}, err
		}
		Likes := Likes{
			CommentId: CommentId,
			PostId:    PostId,
			Value:     value,
		}

		likes = append(likes, Likes)
	}

	query = `
		SELECT userId, comment_id, value
		FROM dislikes
		WHERE postId = ?
	`
	rows, err = db.Query(query, postId)
	if err != nil {
		fmt.Println("error like.go", err)
		return LikeData{}, err
	}
	dislikes := make([]Dislikes, 0)
	for rows.Next() {
		var PostId, CommentId, value int
		err := rows.Scan(&PostId, &CommentId, &value)
		if err != nil {
			return LikeData{}, err
		}
		Dislikes := Dislikes{
			CommentId: CommentId,
			PostId:    PostId,
			Value:     value,
		}
		dislikes = append(dislikes, Dislikes)
	}

	likesData = LikeData{
		Likes:    likes,
		Dislikes: dislikes,
	}

	return likesData, nil
}

func GetUserLikedPosts(userId int) (LikeData, error) {
	var likesData LikeData

	query := `
		SELECT postId, comment_id, value
		FROM likes
		WHERE userId = ?
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		fmt.Println("error like.go", err)
		return LikeData{}, err
	}
	likes := []Likes{}
	for rows.Next() {
		var PostId, CommentId, value int
		err := rows.Scan(&PostId, &CommentId, &value)
		if err != nil {
			return LikeData{}, err
		}
		Likes := Likes{
			CommentId: CommentId,
			PostId:    PostId,
			Value:     value,
		}

		likes = append(likes, Likes)
	}

	query = `
		SELECT postId, comment_id, value
		FROM dislikes
		WHERE userId = ?
	`
	rows, err = db.Query(query, userId)
	if err != nil {
		fmt.Println("error like.go", err)
		return LikeData{}, err
	}
	dislikes := make([]Dislikes, 0)
	for rows.Next() {
		var PostId, CommentId, value int
		err := rows.Scan(&PostId, &CommentId, &value)
		if err != nil {
			return LikeData{}, err
		}
		Dislikes := Dislikes{
			CommentId: CommentId,
			PostId:    PostId,
			Value:     value,
		}
		dislikes = append(dislikes, Dislikes)
	}

	likesData = LikeData{
		Likes:    likes,
		Dislikes: dislikes,
	}

	return likesData, nil
}
func SaveLike(commentId, postId, userId int) {
	// RemoveDislike(commentId, postId, userId)
	_, err := db.Exec("INSERT INTO likes (postId, userId, comment_id, value) VALUES (?, ?, ?, ?)", postId, userId, commentId, 1)
	if err != nil {
		fmt.Println("like.go error inserting like: ", err)
	}
}

func SaveDislike(commentId, postId, userId int) {
	// RemoveLike(commentId, postId, userId)
	_, err := db.Exec("INSERT INTO dislikes (postId, userId, comment_id, value) VALUES (?, ?, ?, ?)", postId, userId, commentId, -1)
	if err != nil {
		fmt.Println("like.go error inserting dislike: ", err)
	}
}

func RemoveLike(commentId, postId, userId int) {
	_, err := db.Exec("DELETE FROM likes WHERE postId = ? AND userId = ? AND comment_id = ? AND value = 1", postId, userId, commentId)
	if err != nil {
		fmt.Println("like.go error removing like: ", err)
	}
}

func RemoveDislike(commentId, postId, userId int) {
	_, err := db.Exec("DELETE FROM dislikes WHERE postId = ? AND userId = ? AND comment_id = ? AND value = -1", postId, userId, commentId)
	if err != nil {
		fmt.Println("like.go error removing dislike: ", err)
	}
}
