package models

import "time"

type Content struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

func getContent(ID int) (Content, error) {
	var content Content

	stmt, err := db.Prepare("SELECT title FROM posts WHERE id = ?")
	if err != nil {
		return Content{}, err
	}
	defer stmt.Close()

	var title string

	err = stmt.QueryRow(ID).Scan(&title)
	if err != nil {
		return Content{}, err
	}

	stmt, err = db.Prepare("SELECT content FROM posts WHERE id = ?")
	if err != nil {
		return Content{}, err
	}
	defer stmt.Close()

	var post string

	err = stmt.QueryRow(ID).Scan(&post)
	if err != nil {
		return Content{}, err
	}

	stmt, err = db.Prepare("SELECT created_at FROM posts WHERE id = ?")
	if err != nil {
		return Content{}, err
	}
	defer stmt.Close()

	var timeCreated time.Time

	err = stmt.QueryRow(ID).Scan(&timeCreated)
	if err != nil {
		return Content{}, err
	}

	content = Content{
		Title: title,
		Content: post,
		CreatedAt: timeCreated,
	}

	return content, nil
}

// func (p *Post) Save() error {
//     // save the post to the database
// }

// func GetPostByID(id int) (*Post, error) {
//     // query the database for a post with the given ID
// }

// func GetPostsByCategoryID(id int) ([]*Post, error) {
//     // query the database for all posts in a given category
// }

// func GetPostsByUserID(id int) ([]*Post, error) {
//     // query the database for all posts created by a given user
// }
