package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserModel struct {
	ID                    uint `gorm:"primaryKey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	Username              string         `gorm:"column:username;unique"`
	Password              string         `gorm:"column:password"`
	LastLoggedInAt        time.Time      `gorm:"column:lastLoggedInAt"`
	LastPasswordUpdatedAt time.Time      `gorm:"column:lastPasswordUpdatedAt"`
}
