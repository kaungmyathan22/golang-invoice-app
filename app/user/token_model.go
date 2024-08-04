package user

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"gorm.io/gorm"
)

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
	p.HashToken()
	return nil
}
