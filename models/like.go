package models

import "fmt"

type LikeData struct {
	PostId int `json:"postId"`
	Value  int `json:"value"`
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

func SaveLike(postId, userId int) {
	_, err := db.Exec("INSERT INTO likes (postId, userId, value) Values (?, ?, ?)", postId, userId, 1)
	if err != nil {
		fmt.Println("Error in post.go, inserting like: ", err)
	}
}

// func getLikes(ID int) (int, error) {
// 	stmt, err := db.Prepare("SELECT value FROM likes WHERE id = ?")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()

// 	var likes int
// 	// Assuming "idValue" is the ID of the comment you want to retrieve
// 	err = stmt.QueryRow(ID).Scan(&likes)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return likes, nil
// }

// func getDislikes(ID int) (int, error) {
// 	stmt, err := db.Prepare("SELECT value FROM likes WHERE id = ?")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()

// 	var dislikes int
// 	// Assuming "idValue" is the ID of the comment you want to retrieve
// 	err = stmt.QueryRow(ID).Scan(&dislikes)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return dislikes, nil
// }
//  {
//     ID     int `json:"id"`
//     PostID int `json:"post_id"`
//     UserID int `json:"user_id"`
// }

// func (l *Like) Save() error {
//     // save the like to the database
// }

// func GetLikesByPostID(id int) ([]*Like, error) {
//     // query the database for all likes on a given post
// }
