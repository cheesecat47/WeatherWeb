package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// User struct
type User struct {
	UserID      string         `db:"user_id"`
	UserEmail   string         `db:"user_email"`
	UserPW      string         `db:"user_pw"`
	UserName    string         `db:"user_name"`
	RawUserRole sql.NullString `db:"user_role"`
	UserRole    string
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

// checkNullString func
func (u *User) checkNullString() {
	log.Println("User_model: checkNullString")

	if u.RawUserRole.Valid {
		u.UserRole = u.RawUserRole.String
	} else {
		u.UserRole = "NULL"
	}
}

// GetUserByID func
func GetUserByID(id string) (User, error) {
	user := User{}
	query := fmt.Sprintf("SELECT * FROM user WHERE user_id = '%s'", id)
	// log.Println("User_model: GetUserByID: query:", query)

	err := DB.QueryRow(query).Scan(
		&user.UserID, &user.UserEmail, &user.UserPW, &user.UserName,
		&user.RawUserRole, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		log.Println("User_model: GetUserByID: err: ", err)
		return user, err
	}

	user.checkNullString()
	return user, nil
}
