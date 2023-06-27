package models

import (
	"fmt"
	"net/http"
)

func SessionIsActive(sessionId string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE sessionId = ?", sessionId).Scan(&count)
	if err != nil {
		fmt.Println("no cookie, ya cant post")
		return false, err
	}
	fmt.Println("valid sesh id")
	return count > 0, nil
}

func SavePost(title,content,category string, userId int){
	_,err := db.Exec("INSERT INTO posts (title, content, category, userId) Values (?, ?, ?, ?)", title, content, category, userId)
	if err != nil{
		fmt.Println("Error inserting into posts: ", err)
		return
	}
	fmt.Println("Successfully inserted into posts!!!!!!!")
}

func GetUserByCookie(r *http.Request) (int, error) {
	cookie, _ := r.Cookie("session")
	sessionId := cookie.Value
	var userId int
	err := db.QueryRow("SELECT id FROM users WHERE sessionId = ?", sessionId).Scan(&userId)
	if err != nil {
		// Handle the database query error accordingly
		fmt.Println("user has no sesh id rn")
		return 0, err
	}
	fmt.Printf("the user id is: %v", userId)
	return userId, nil
}

// type Content struct {
// 	ID         int       `json:"id"`
// 	Title      string    `json:"title"`
// 	Content    string    `json:"content"`
// 	CreatedAt  time.Time `json:"created_at"`
// }

// func getUserContent(ID int) (Content, error) {
// 	var content Content

// 	stmt, err := db.Prepare("SELECT title FROM posts WHERE id = ?")
// 	if err != nil {
// 		return Content{}, err
// 	}
// 	defer stmt.Close()

// 	var title string

// 	err = stmt.QueryRow(ID).Scan(&title)
// 	if err != nil {
// 		return Content{}, err
// 	}

// 	stmt, err = db.Prepare("SELECT content FROM posts WHERE id = ?")
// 	if err != nil {
// 		return Content{}, err
// 	}
// 	defer stmt.Close()

// 	var post string

// 	err = stmt.QueryRow(ID).Scan(&post)
// 	if err != nil {
// 		return Content{}, err
// 	}

// 	stmt, err = db.Prepare("SELECT created_at FROM posts WHERE id = ?")
// 	if err != nil {
// 		return Content{}, err
// 	}
// 	defer stmt.Close()

// 	var timeCreated time.Time

// 	err = stmt.QueryRow(ID).Scan(&timeCreated)
// 	if err != nil {
// 		return Content{}, err
// 	}

// 	content = Content{
// 		Title: title,
// 		Content: post,
// 		CreatedAt: timeCreated,
// 	}

// 	return content, nil
// }

// func GetAllContent() (Content, error){

// 	return Content{}, nil
// }
