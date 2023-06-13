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
			category_id INTEGER,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(category_id) REFERENCES categories(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY,
			content TEXT,
			post_id INTEGER,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(post_id) REFERENCES posts(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS likes (
			id INTEGER PRIMARY KEY,
			post_id INTEGER,
			comment_id INTEGER,
			user_id INTEGER,
			value INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(post_id) REFERENCES posts(id),
			FOREIGN KEY(comment_id) REFERENCES comments(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		);

        CREATE TABLE IF NOT EXISTS dislikes (
            id INTEGER PRIMARY KEY,
            post_id INTEGER,
            comment_id INTEGER,
            user_id INTEGER,
            value INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
            FOREIGN KEY (comment_id) REFERENCES comments (id) ON DELETE CASCADE,
            FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
        );
        
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Set database connection pool limits
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60 * 1000)

	fmt.Println("Database initialized")
}

func InsertDB(username, email, password string) error {
	// Prepare the SQL statement to insert a new post
	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Execute the prepared statement with the provided values and current timestamp
	_, err = stmt.Exec(username, email, password)
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

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed")
	}
}
