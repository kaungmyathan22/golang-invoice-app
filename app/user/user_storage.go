package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserStorage interface {
	GetCount(condition interface{}) (int64, error)
	GetAll(page, pageSize int) ([]UserModel, error)
	GetById(id uint) (*UserModel, error)
	GetByUsername(username string) (*UserModel, error)
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

func (storage *UserStorageImpl) GetCount(condition interface{}) (int64, error) {
	var totalRecords int64
	query := storage.db.Model(&UserModel{})
	if condition != nil {
		query = query.Where(condition)
	}
	result := query.Count(&totalRecords)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalRecords, nil
}

func (storage *UserStorageImpl) GetAll(page, pageSize int) ([]UserModel, error) {
	var users []UserModel
	offset := (page - 1) * pageSize
	result := storage.db.Offset(offset).Limit(pageSize).Find(&users)
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

func (storage *UserStorageImpl) GetByUsername(username string) (*UserModel, error) {
	var user *UserModel
	result := storage.db.Where("username = ?", username).First(&user)
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
