package models

import (
	"time"

	"gorm.io/gorm"

	db "github.com/nomannaq/webshop-api-golang/cmd/database"
)

type User struct {
	gorm.Model
	Username  string         `gorm:"unique;not null" json:"username"`
	Email     string         `gorm:"unique;not null" json:"email"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Role      string         `gorm:"not null, default:user" json:"role"`
}

func CreateUser(user *User) error {
	return db.DB.Create(user).Error
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
