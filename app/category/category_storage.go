package category

import (
	"errors"

	"gorm.io/gorm"
)

type CategoryStorage interface {
	GetCount(condition interface{}) (int64, error)
	GetAll(page, pageSize int) ([]CategoryModel, error)
	GetById(id uint) (*CategoryModel, error)
	Create(category CategoryModel) (*CategoryModel, error)
	Update(category CategoryModel) error
	Delete(id uint) error
}

type CategoryStorageImpl struct {
	db *gorm.DB
}

func NewCategoryStorage(db *gorm.DB) *CategoryStorageImpl {
	return &CategoryStorageImpl{db: db}
}

func (storage *CategoryStorageImpl) GetCount(condition interface{}) (int64, error) {
	var totalRecords int64
	query := storage.db.Model(&CategoryModel{})
	if condition != nil {
		query = query.Where(condition)
	}
	result := query.Count(&totalRecords)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalRecords, nil
}

func (storage *CategoryStorageImpl) GetAll(page, pageSize int) ([]CategoryModel, error) {
	var categorys []CategoryModel
	offset := (page - 1) * pageSize
	result := storage.db.Offset(offset).Limit(pageSize).Find(&categorys)
	if result.Error != nil {
		return nil, result.Error
	}
	return categorys, nil
}

func (storage *CategoryStorageImpl) GetById(id uint) (*CategoryModel, error) {
	var category *CategoryModel
	result := storage.db.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, result.Error
	}
	return category, nil
}

func (storage *CategoryStorageImpl) GetByCategoryname(categoryname string) (*CategoryModel, error) {
	var category *CategoryModel
	result := storage.db.Where("name = ?", categoryname).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, result.Error
	}
	return category, nil
}

func (storage *CategoryStorageImpl) Create(category CategoryModel) (*CategoryModel, error) {
	result := storage.db.Create(&category)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (storage *CategoryStorageImpl) Update(category CategoryModel) error {
	result := storage.db.Save(category)
	return result.Error
}

func (storage *CategoryStorageImpl) Delete(id uint) error {
	result := storage.db.Delete(&CategoryModel{}, id)
	return result.Error
}
