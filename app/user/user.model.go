package user

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("email already exists")
)

type UserModel struct {
	ID                    uint `gorm:"primaryKey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	Username              string         `gorm:"column:username"`
	Email                 string         `gorm:"type:citext;column:email;unique;not null"`
	Password              string         `gorm:"column:password;not null"`
	LastLoggedInAt        time.Time      `gorm:"column:lastLoggedInAt"`
	LastPasswordUpdatedAt time.Time      `gorm:"column:lastPasswordUpdatedAt"`
}

func (UserModel) TableName() string {
	return "users"
}

type PasswordResetTokenModel struct {
	gorm.Model
	TokenHash string `gorm:"uniqueIndex;not null"`
	UserID    uint   `gorm:"not null"`
	User      UserModel
	ExpiresAt time.Time `gorm:"not null"`
}

func (PasswordResetTokenModel) TableName() string {
	return "password_reset"
}

func (p *PasswordResetTokenModel) HashToken() {
	hash := sha256.New()
	hash.Write([]byte(p.TokenHash))
	p.TokenHash = hex.EncodeToString(hash.Sum(nil))
}

func (p *PasswordResetTokenModel) GenerateToken() error {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}
	p.TokenHash = hex.EncodeToString(b)
	return nil
}
