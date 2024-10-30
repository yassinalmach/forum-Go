package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DataBase *sql.DB

func InitDb() error {
	// Initialize database connection
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}

	db.Ping()

	_, err = db.Exec(`-- Users table
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
);
-- Sessions
CREATE TABLE IF NOT EXISTS sessions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	session_id TEXT NOT NULL UNIQUE,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
-- Posts table
CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	image TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
-- Comments table
CREATE TABLE IF NOT EXISTS comments (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	content TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (post_id) REFERENCES posts(id)
);
-- Likes table
CREATE TABLE IF NOT EXISTS engagement (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	is_like BOOLEAN ,
	is_dislike BOOLEAN,
	user_id INTEGER NOT NULL,
	post_id INTEGER,
	comment_id INTEGER,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (post_id) REFERENCES posts(id),
	FOREIGN KEY (comment_id) REFERENCES comments(id),
	CHECK (
		(
			post_id IS NOT NULL
			AND comment_id IS NULL
		)
		OR (
			post_id IS NULL
			AND comment_id IS NOT NULL
		)
	)
);
-- Likes table	
CREATE TABLE IF NOT EXISTS categories(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	categori TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (post_id) REFERENCES posts(id)
);`)
	if err != nil {
		return err
	}

	DataBase = db

	return err
}
