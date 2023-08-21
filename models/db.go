package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tables if they don't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			email TEXT UNIQUE,
			username TEXT UNIQUE,
			password TEXT,
			sessionId TEXT UNIQUE
		);

		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY,
			name TEXT UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY,
			title TEXT,
			content TEXT,
			userId INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(userId) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS post_categories (
			id INTEGER PRIMARY KEY,
			post_id INTEGER,
			category_id INTEGER,
			FOREIGN KEY (post_id) REFERENCES posts(id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);

		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY,
			content TEXT,
			postId INTEGER,
			userId INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(postId) REFERENCES posts(id),
			FOREIGN KEY(userId) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS likes (
			id INTEGER PRIMARY KEY,
			postId INTEGER,
			comment_id INTEGER,
			userId INTEGER,
			value INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(postId) REFERENCES posts(id),
			FOREIGN KEY(comment_id) REFERENCES comments(id),
			FOREIGN KEY(userId) REFERENCES users(id)
		);

        CREATE TABLE IF NOT EXISTS dislikes (
            id INTEGER PRIMARY KEY,
            postId INTEGER,
            comment_id INTEGER,
            userId INTEGER,
            value INTEGER,
            FOREIGN KEY (postId) REFERENCES posts (id) ON DELETE CASCADE,
            FOREIGN KEY (comment_id) REFERENCES comments (id) ON DELETE CASCADE,
            FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE
        );
        
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err1 := db.Exec("INSERT OR IGNORE INTO categories (name) VALUES (?)", "etymology")
	if err1 != nil {
		fmt.Println("cant insert into categoriy at the start")
	}
	_, err2 := db.Exec("INSERT OR IGNORE INTO categories (name) VALUES (?)", "biology")
	if err2 != nil {
		fmt.Println("cant insert into categoriy at the start")
	}

	// Set database connection pool limits
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60 * 1000)

	fmt.Printf("Database initialized\n")
}

func InsertDB(username, email, password, seshId string) error {
	// Prepare the SQL statement to insert a new post
	stmt, err := db.Prepare("INSERT INTO users (username, email, password, sessionId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Execute the prepared statement with the provided values and current timestamp
	_, err = stmt.Exec(username, email, password, seshId)
	if err != nil {
		return err
	}

	return nil
}

func GetID(email string) (int, error) {
	stmt, err := db.Prepare("SELECT id FROM users WHERE email = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var userID int
	// Assuming "emailValue" is the email you want to search for
	err = stmt.QueryRow(email).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// func DeleteLikesTable() {
// 	_, err := db.Exec("DELETE FROM dislikes")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed")
	}
}
