package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-" means this won't be included in JSON
}

type Post struct {
	Id         int        `json:"id"`
	UserID     int        `json:"user_id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	User       *User      `json:"user,omitempty"`
	Categories []Category `json:"categories,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	Likes      int        `json:"likes"`
	Dislikes   int        `json:"dislikes"`
}

type Comment struct {
	Id        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	User      *User     `json:"user,omitempty"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Like struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	PostID    *int64    `json:"post_id,omitempty"`
	CommentID *int64    `json:"comment_id,omitempty"`
	IsLike    bool      `json:"is_like"`
	CreatedAt time.Time `json:"created_at"`
}

// User methods
func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

type Session struct {
	ID        string    `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
