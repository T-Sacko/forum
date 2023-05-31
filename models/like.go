package models



func getLikes(ID int) (int, error) {
	stmt, err := db.Prepare("SELECT value FROM likes WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var likes int
	// Assuming "idValue" is the ID of the comment you want to retrieve
	err = stmt.QueryRow(ID).Scan(&likes)
	if err != nil {
		return 0, err
	}
	return likes, nil
}

func getDislikes(ID int) (int, error) {
	stmt, err := db.Prepare("SELECT value FROM likes WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var dislikes int
	// Assuming "idValue" is the ID of the comment you want to retrieve
	err = stmt.QueryRow(ID).Scan(&dislikes)
	if err != nil {
		return 0, err
	}
	return dislikes, nil
}
// type Like struct {
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
