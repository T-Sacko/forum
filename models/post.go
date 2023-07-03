package models

import (
	"fmt"
	"net/http"
)

type Post struct {
	ID       int
	Title      string
	Content    string
	Categories []string
}

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

func SavePost(title, content string, userId int) int {
	result, err := db.Exec("INSERT INTO posts (title, content, userId) Values (?, ?, ?)", title, content, userId)
	if err != nil {
		fmt.Println("Error inserting into posts: ", err)
		return 0
	}
	fmt.Println("Successfully inserted into posts!!!!!!!")
	postId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("error with getting postid from lastInserId")
	}
	return int(postId)
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
	fmt.Printf("the user id is: %v\n", userId)
	return userId, nil
}

func getPostsFromDB() ([]Post, error) {
	query := `
		SELECT posts.id, posts.title, posts.content, categories.name
		FROM posts
		INNER JOIN post_categories ON posts.id = post_categories.post_id
		INNER JOIN categories ON post_categories.category_id = categories.id
		ORDER BY posts.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	currentPost := Post{}
	for rows.Next() {
		var postID int
		var title, content, categoryName string
		err := rows.Scan(&postID, &title, &content, &categoryName)
		if err != nil {
			return nil, err
		}

		if currentPost.ID != postID {
			if currentPost.ID != 0 {
				posts = append(posts, currentPost)
			}
			currentPost = Post{
				ID:      postID,
				Title:   title,
				Content: content,
			}
		}

		currentPost.Categories = append(currentPost.Categories, categoryName)
	}

	if currentPost.ID != 0 {
		posts = append(posts, currentPost)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
