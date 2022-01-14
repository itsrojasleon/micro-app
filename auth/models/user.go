package models

import (
	"gorm.io/gorm"

	"github.com/rojasleon/reserve-micro/auth/internal"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Hash user's password before creating it
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword := internal.HashAndSalt([]byte(u.Password))
	u.Password = hashedPassword

	return
}
