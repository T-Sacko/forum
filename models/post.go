package models

import (
	"fmt"
	"net/http"
	"strings"
)

type Post struct {
	ID         int
	Title      string
	Content    string
	Username   string
	Categories []string
	Likes      int
	Dislikes   int
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

func GetPostsFromDB() ([]Post, error) {
	query := `
		SELECT posts.id, posts.title, posts.content, users.username,
			GROUP_CONCAT(DISTINCT categories.name) AS categoryNames,
			COALESCE(SUM(CASE WHEN likes.value = 1 THEN 1 ELSE 0 END), 0) AS likes,
			COALESCE(SUM(CASE WHEN likes.value = -1 THEN 1 ELSE 0 END), 0) AS dislikes
		FROM posts
		INNER JOIN users ON posts.userId = users.id
		INNER JOIN post_categories ON posts.id = post_categories.post_id
		INNER JOIN categories ON post_categories.category_id = categories.id
		LEFT JOIN likes ON likes.postId = posts.id
		GROUP BY posts.id
		ORDER BY posts.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var postID int
		var title, content, username, categoryNames string
		var likes, dislikes int
		err := rows.Scan(&postID, &title, &content, &username, &categoryNames, &likes, &dislikes)
		if err != nil {
			return nil, err
		}

		categories := strings.Split(categoryNames, ",")
		post := Post{
			ID:         postID,
			Title:      title,
			Content:    content,
			Username:   username,
			Categories: categories,
			Likes:      likes,
			Dislikes:   dislikes,
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// func findPostByID(posts []Post, postID int) *Post {
// 	for i := range posts {
// 		if posts[i].ID == postID {
// 			return &posts[i]
// 		}
// 	}
// 	return nil
// }

// func containsCategory(categories []string, category string) bool {
// 	for _, c := range categories {
// 		if c == category {
// 			return true
// 		}
// 	}
// 	return false
// }
