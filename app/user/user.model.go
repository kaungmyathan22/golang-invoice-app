package user

import (
	"errors"
	"time"

	"github.com/kaungmyathan22/golang-invoice-app/app/common"
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

type UserStorage interface {
	GetAll(payload common.PaginationParamsRequest) ([]UserModel, error)
	GetById(id uint) (*UserModel, error)
	Create(user UserModel) (*UserModel, error)
	Update(user UserModel) error
	Delete(id int) error
}

type UserStorageImpl struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorageImpl {
	return &UserStorageImpl{db: db}
}

func (storage *UserStorageImpl) GetAll(payload common.PaginationParamsRequest) ([]UserModel, error) {
	var users []UserModel
	offset := (payload.Page - 1) * payload.PageSize
	result := storage.db.Offset(offset).Limit(payload.PageSize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (storage *UserStorageImpl) GetById(id uint) (*UserModel, error) {
	var user *UserModel
	result := storage.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}
	return user, nil
}

func (storage *UserStorageImpl) Create(user UserModel) (*UserModel, error) {
	result := storage.db.Create(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (storage *UserStorageImpl) Update(user UserModel) error {
	result := storage.db.Save(user)
	return result.Error
}

func (storage *UserStorageImpl) Delete(id int) error {
	result := storage.db.Delete(&UserModel{}, id)
	return result.Error
}
