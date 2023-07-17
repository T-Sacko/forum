package models

import (
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	ID         int
	Title      string
	Content    string
	Username   string
	Categories []string
	Comments   []Comment
	Likes      int
	Dislikes   int
}

func SessionIsActive(sessionId string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE sessionId = ?", sessionId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func SavePost(title, content string, userId int) (int, error) {
	result, err := db.Exec("INSERT INTO posts (title, content, userId) Values (?, ?, ?)", title, content, userId)
	if err != nil {
		return 0, err
	}
	postId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(postId), nil
}

func GetUserByCookie(r *http.Request) (*User, error) {

	var username string

	cookie, err := r.Cookie("session")
	if err != nil {
		return &User{}, err
	}
	sessionId := cookie.Value
	var userId int
	err = db.QueryRow("SELECT id, username FROM users WHERE sessionId = ?", sessionId).Scan(&userId, &username)
	if err != nil {
		return &User{}, err
	}

	user := &User{ID: userId, Username: username}

	return user, nil
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
		GROUP BY posts.id, users.username
		ORDER BY posts.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var title, content, username, categoryNames string
		var postID int
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

		comments, err := getCommentsForPost(postID)
		if err != nil {
			return nil, err
		}
		post.Comments = comments

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	posts = reversePosts(posts)

	return posts, nil
}

func getCommentsForPost(postID int) ([]Comment, error) {
	query := `
		SELECT comments.content, comments.userId
		FROM comments
		WHERE comments.postId = ?
		ORDER BY comments.id
	`

	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var content string
		var userID int
		err := rows.Scan(&content, &userID)
		if err != nil {
			return nil, err
		}
		comment := Comment{
			UserID:  userID,
			PostID:  strconv.Itoa(postID),
			Comment: content,
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func reversePosts(posts []Post) []Post {
	var reversedPosts []Post
	for i := len(posts) - 1; i >= 0; i-- {
		reversedPosts = append(reversedPosts, posts[i])
	}
	return reversedPosts
}
