package model

import (
	"time"
)

// Board struct
type Board struct {
	BoardID   int        `db:"board_id"`
	BoardName string     `db:"board_name"`
	UserID    string     `db:"user_id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
