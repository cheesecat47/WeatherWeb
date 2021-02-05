package model

import (
	"time"
)

// Token struct
type Token struct {
	UserID                string     `db:"user_id"`
	RefreshToken          string     `db:"refresh_token"`
	RefreshTokenExpiredAt *time.Time `db:"refresh_token_expired_at"`
	FcmToken              string     `db:"fcm_token"`
	CreatedAt             *time.Time `db:"created_at"`
	UpdatedAt             *time.Time `db:"updated_at"`
	DeletedAt             *time.Time `db:"deleted_at"`
}
