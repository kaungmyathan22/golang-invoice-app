package user_models

import (
	"errors"
	"time"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type UserModel struct {
	ID                    uint `gorm:"primaryKey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	UserName              string         `gorm:"column:username"`
	Password              string         `gorm:"column:password"`
	LastLoggedInAt        time.Time      `gorm:"column:lastLoggedInAt"`
	LastPasswordUpdatedAt time.Time      `gorm:"column:lastPasswordUpdatedAt"`
}

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (storage *UserStorage) GetAll() ([]UserModel, error) {
	var users []UserModel
	result := storage.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (storage *UserStorage) GetById(id uint) (*UserModel, error) {
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

func (s UserStorage) Create(user UserModel) (error, *UserModel) {
	result := s.db.Create(&user)
	if err := result.Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrEmailAlreadyExists, nil
		}
		return err, nil
	}
	return nil, &user
}

func (s UserStorage) Update(user UserModel) error {
	result := s.db.Save(user)
	return result.Error
}

func (s UserStorage) Delete(id int) error {
	result := s.db.Delete(&UserModel{}, id)
	return result.Error
}
