package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Post struct {
	ID             int
	Title          string
	Content        string
	Username       string
	Categories     []string
	FileName       string
	FileType       string
	Likes          int
	Dislikes       int
	Comments       []Comment
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserLikeStatus int
}

func SessionIsActive(sessionId string) (int, bool, error) {
	var userId int
	err := db.QueryRow("SELECT id FROM users WHERE sessionId = ?", sessionId).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No cookie, you can't post")
			return 0, false, err
		}
		fmt.Println("Error retrieving session ID:", err)
		return 0, false, err
	}
	return userId, true, nil
}

func SavePost(title, content, fileName string, userId int) (int, error) {
	var result sql.Result
	var err error

	if fileName != "" {
		result, err = db.Exec("INSERT INTO posts (title, content, userId, fileName) Values (?, ?, ?, ?)", title, content, userId, fileName)
	} else {
		result, err = db.Exec("INSERT INTO posts (title, content, userId) Values (?, ?, ?)", title, content, userId)
	}

	if err != nil {
		fmt.Println("Error inserting into posts: ", err)
		return 0, err
	}

	postId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("error with getting postid from lastInsertId")
		return 0, err
	}

	return int(postId), nil
}

var fileName sql.NullString

func GetPostsFromDB(userID int) ([]Post, error) {
	query := `
        SELECT 
            posts.id, posts.title, posts.content, users.username, posts.fileName,
            COALESCE(GROUP_CONCAT(DISTINCT categories.name), '') AS categoryNames,
            COALESCE(SUM(CASE WHEN likes.value = 1 THEN 1 ELSE 0 END), 0) AS likes,
            COALESCE(SUM(CASE WHEN likes.value = -1 THEN 1 ELSE 0 END), 0) AS dislikes,
            COALESCE(likes_for_user.value, 0) AS userLikeStatus
        FROM 
            posts
        INNER JOIN 
            users ON posts.userId = users.id
        LEFT JOIN 
            post_categories ON posts.id = post_categories.post_id
        LEFT JOIN 
            categories ON post_categories.category_id = categories.id
        LEFT JOIN 
            likes ON likes.postId = posts.id
        LEFT JOIN 
            likes AS likes_for_user ON likes_for_user.postId = posts.id AND likes_for_user.userId = ?
        GROUP BY 
            posts.id    
        ORDER BY 
            posts.id DESC
    `

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var postID int
		var title, content, username, categoryNames string
		var likes, dislikes, userLikeStatus int
		err := rows.Scan(&postID, &title, &content, &username, &fileName, &categoryNames, &likes, &dislikes, &userLikeStatus)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		var categories []string
		if categoryNames != "" {
			categories = strings.Split(categoryNames, ",")
		}

		fileName1:= ""
		if fileName.Valid {
			fileName1 = fileName.String
		} else {
			// Handle the case where fileName is NULL (no file uploaded)
			fileName1 = "" // Or any other appropriate value
		}

		post := Post{
			ID:             postID,
			Title:          title,
			Content:        content,
			Username:       username,
			FileName:       fileName1,
			Categories:     categories,
			Likes:          likes,
			Dislikes:       dislikes,
			UserLikeStatus: userLikeStatus,
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func FilterByCategory(userID int, category string) ([]Post, error) {
	query := `
        SELECT 
            posts.id, posts.title, posts.content, users.username,
            COALESCE(GROUP_CONCAT(DISTINCT categories.name), '') AS categoryNames,
            COALESCE(SUM(CASE WHEN likes.value = 1 THEN 1 ELSE 0 END), 0) AS likes,
            COALESCE(SUM(CASE WHEN likes.value = -1 THEN 1 ELSE 0 END), 0) AS dislikes,
            COALESCE(likes_for_user.value, 0) AS userLikeStatus
        FROM 
            posts
        INNER JOIN 
            users ON posts.userId = users.id
        LEFT JOIN 
            post_categories ON posts.id = post_categories.post_id
        LEFT JOIN 
            categories ON post_categories.category_id = categories.id
        LEFT JOIN 
            likes ON posts.id = likes.postId
        LEFT JOIN 
            likes AS likes_for_user ON likes_for_user.postId = posts.id AND likes_for_user.userId = ?
        WHERE 
            categories.name = ?
        GROUP BY 
            posts.id
        ORDER BY 
            posts.id DESC
    `

	rows, err := db.Query(query, userID, category)
	if err != nil {
		fmt.Println("error inna a ting")
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var postID int
		var title, content, username, categoryNames string
		var likes, dislikes, userLikeStatus int
		err := rows.Scan(&postID, &title, &content, &username, &categoryNames, &likes, &dislikes, &userLikeStatus)
		if err != nil {
			return nil, err
		}

		categories := strings.Split(categoryNames, ",")
		fileName1:= ""
		if fileName.Valid {
			fileName1 = fileName.String
		} else {
			// Handle the case where fileName is NULL (no file uploaded)
			fileName1 = "" // Or any other appropriate value
		}

		post := Post{
			ID:             postID,
			Title:          title,
			Content:        content,
			Username:       username,
			FileName:       fileName1,
			Categories:     categories,
			Likes:          likes,
			Dislikes:       dislikes,
			UserLikeStatus: userLikeStatus,
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func FilterByLiked(userID int) ([]Post, error) {
	query := `
		SELECT 
			posts.id, posts.title, posts.content, users.username,
			COALESCE(GROUP_CONCAT(DISTINCT categories.name), '') AS categoryNames,
			COALESCE(SUM(CASE WHEN likes.value = 1 THEN 1 ELSE 0 END), 0) AS likes,
			COALESCE(SUM(CASE WHEN likes.value = -1 THEN 1 ELSE 0 END), 0) AS dislikes,
			COALESCE(likes_for_user.value, 0) AS userLikeStatus
		FROM 
			posts
		INNER JOIN 
			users ON posts.userId = users.id
		LEFT JOIN 
			post_categories ON posts.id = post_categories.post_id
		LEFT JOIN 
			categories ON post_categories.category_id = categories.id
		LEFT JOIN 
			likes ON posts.id = likes.postId
		LEFT JOIN 
			likes AS likes_for_user ON likes_for_user.postId = posts.id AND likes_for_user.userId = ?
		GROUP BY 
			posts.id
		HAVING 
			likes > 0 AND EXISTS (
				SELECT 1 FROM likes WHERE likes.postId = posts.id AND likes.userId = ? AND likes.value = 1
			)
		ORDER BY 
			posts.id DESC
	`

	rows, err := db.Query(query, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var postID int
		var title, content, username, categoryNames string
		var likes, dislikes, userLikeStatus int
		err := rows.Scan(&postID, &title, &content, &username, &categoryNames, &likes, &dislikes, &userLikeStatus)
		if err != nil {
			return nil, err
		}

		categories := strings.Split(categoryNames, ",")
		fileName1:= ""
		if fileName.Valid {
			fileName1 = fileName.String
		} else {
			// Handle the case where fileName is NULL (no file uploaded)
			fileName1 = "" // Or any other appropriate value
		}

		post := Post{
			ID:             postID,
			Title:          title,
			Content:        content,
			Username:       username,
			FileName:       fileName1,
			Categories:     categories,
			Likes:          likes,
			Dislikes:       dislikes,
			UserLikeStatus: userLikeStatus,
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func FilterByUserPosts(userID int) ([]Post, error) {
	query := `
		SELECT
			posts.id, posts.title, posts.content, users.username, 
			COALESCE(GROUP_CONCAT(DISTINCT categories.name), '') AS categoryNames,
			COALESCE(SUM(CASE WHEN likes.value = 1 THEN 1 ELSE 0 END), 0) AS likes,
			COALESCE(SUM(CASE WHEN likes.value = -1 THEN 1 ELSE 0 END), 0) AS dislikes,
			COALESCE(likes_for_user.value, 0) AS userLikeStatus
		FROM 
			posts
		INNER JOIN 
			users ON posts.userId = users.id
		LEFT JOIN 
			post_categories ON posts.id = post_categories.post_id
		LEFT JOIN 
			categories ON post_categories.category_id = categories.id
		LEFT JOIN 
			likes ON likes.postId = posts.id
		LEFT JOIN 
			likes AS likes_for_user ON likes_for_user.postId = posts.id AND likes_for_user.userId = ?
		WHERE 
			users.id = ?
		GROUP BY 
			posts.id
		ORDER BY 
			posts.id DESC
	`

	rows, err := db.Query(query, userID, userID)
	if err != nil {
		fmt.Println(err, "errhere")
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var postID int
		var title, content, username, categoryNames string
		var likes, dislikes, userLikeStatus int
		err := rows.Scan(&postID, &title, &content, &username, &categoryNames, &likes, &dislikes, &userLikeStatus)
		if err != nil {
			fmt.Println(err, "errhere")

			return nil, err
		}

		categories := strings.Split(categoryNames, ",")
		fileName1:= ""
		if fileName.Valid {
			fileName1 = fileName.String
		} else {
			// Handle the case where fileName is NULL (no file uploaded)
			fileName1 = "" // Or any other appropriate value
		}

		post := Post{
			ID:             postID,
			Title:          title,
			Content:        content,
			Username:       username,
			FileName:       fileName1,
			Categories:     categories,
			Likes:          likes,
			Dislikes:       dislikes,
			UserLikeStatus: userLikeStatus,
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err, "errhere")

		return nil, err
	}

	return posts, nil
}
