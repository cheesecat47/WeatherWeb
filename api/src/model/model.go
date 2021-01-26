package model

import (
	"time"

	"gorm.io/gorm"
)

// User struct
type User struct {
	UserID    string `gorm:"primaryKey"`
	UserEmail string
	UserPW    string
	UserName  string
	UserRole  string    `gorm:"default:NULL"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

// Article struct
type Article struct {
	ArticleID      int `gorm:"primaryKey"`
	BoardID        int
	UserID         string
	ArticleTitle   string
	ArticleContent string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

// Board struct
type Board struct {
	BoardID   int `gorm:"primaryKey"`
	BoardName string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Token struct
type Token struct {
	UserID                string `gorm:"primaryKey"`
	RefreshToken          string `gorm:"unique"`
	RefreshTokenExpiredAt time.Time
	FcmToken              string `gorm:"unique"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}

// https://gorm.io/ko_KR/docs/models.html#embedded_struct
