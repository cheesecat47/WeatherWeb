package model

import (
	"time"
)

// Article struct
type Article struct {
	ArticleID      int       `db:"article_id"`
	BoardID        int       `db:"board_id"`
	UserID         string    `db:"user_id"`
	ArticleTitle   string    `db:"article_title"`
	ArticleContent string    `db:"article_content"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at"`
}
