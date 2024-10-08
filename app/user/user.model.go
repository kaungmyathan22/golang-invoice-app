package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("email already exists")
)

type UserModel struct {
	CreatedAt             time.Time
	UpdatedAt             time.Time
	LastLoggedInAt        time.Time      `gorm:"column:lastLoggedInAt"`
	LastPasswordUpdatedAt time.Time      `gorm:"column:lastPasswordUpdatedAt"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	Username              string         `gorm:"column:username"`
	Email                 string         `gorm:"type:citext;column:email;unique;not null"`
	Password              string         `gorm:"column:password;not null"`
	ID                    uint           `gorm:"primaryKey"`
}

func (UserModel) TableName() string {
	return "users"
}
