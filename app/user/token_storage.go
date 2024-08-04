package user

import "gorm.io/gorm"

type TokenStorage interface {
	GetByHash(id uint) (*UserModel, error)
	Create(token PasswordResetTokenModel) (*UserModel, error)
	Delete(id uint) error
}

type TokenStorageImpl struct {
	db *gorm.DB
}

func NewTokenStorage(db *gorm.DB) *TokenStorageImpl {
	return &TokenStorageImpl{db: db}
}
