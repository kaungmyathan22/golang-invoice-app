package user

import (
	"gorm.io/gorm"
)

type TokenStorage interface {
	GetByHash(hash string) (*PasswordResetTokenModel, error)
	Create(token PasswordResetTokenModel) (*PasswordResetTokenModel, error)
	Delete(id uint) error
	DeleteByUserId(userId uint) error
}

type TokenStorageImpl struct {
	db *gorm.DB
}

func NewTokenStorage(db *gorm.DB) *TokenStorageImpl {
	return &TokenStorageImpl{db: db}
}

func (storage *TokenStorageImpl) Create(token PasswordResetTokenModel) (*PasswordResetTokenModel, error) {
	result := storage.db.Create(&token)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (storage *TokenStorageImpl) GetByHash(hash string) (*PasswordResetTokenModel, error) {
	var token *PasswordResetTokenModel
	result := storage.db.Where("token_hash = ?", hash).First(&token)
	if result.Error != nil {
		return nil, result.Error
	}
	return token, nil
}

func (storage *TokenStorageImpl) Delete(id uint) error {
	result := storage.db.Delete(&PasswordResetTokenModel{}, id)
	return result.Error
}

func (storage *TokenStorageImpl) DeleteByUserId(userId uint) error {
	result := storage.db.Where("user_id = ?", userId).Delete(&PasswordResetTokenModel{})
	return result.Error
}
