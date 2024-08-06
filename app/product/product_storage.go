package product

import (
	"errors"

	"gorm.io/gorm"
)

type ProductStorage interface {
	GetCount(condition interface{}) (int64, error)
	GetAll(page, pageSize int) ([]ProductModel, error)
	GetById(id uint) (*ProductModel, error)
	Create(Product ProductModel) (*ProductModel, error)
	Update(Product ProductModel) error
	Delete(id uint) error
}

type ProductStorageImpl struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) *ProductStorageImpl {
	return &ProductStorageImpl{db: db}
}

func (storage *ProductStorageImpl) GetCount(condition interface{}) (int64, error) {
	var totalRecords int64
	query := storage.db.Model(&ProductModel{})
	if condition != nil {
		query = query.Where(condition)
	}
	result := query.Count(&totalRecords)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalRecords, nil
}

func (storage *ProductStorageImpl) GetAll(page, pageSize int) ([]ProductModel, error) {
	var Products []ProductModel
	offset := (page - 1) * pageSize
	result := storage.db.Offset(offset).Limit(pageSize).Find(&Products)
	if result.Error != nil {
		return nil, result.Error
	}
	return Products, nil
}

func (storage *ProductStorageImpl) GetById(id uint) (*ProductModel, error) {
	var Product *ProductModel
	result := storage.db.First(&Product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, result.Error
	}
	return Product, nil
}

func (storage *ProductStorageImpl) GetByProductname(Productname string) (*ProductModel, error) {
	var Product *ProductModel
	result := storage.db.Where("name = ?", Productname).First(&Product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, result.Error
	}
	return Product, nil
}

func (storage *ProductStorageImpl) Create(product ProductModel) (*ProductModel, error) {
	result := storage.db.Create(&product)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (storage *ProductStorageImpl) Update(Product ProductModel) error {
	result := storage.db.Save(Product)
	return result.Error
}

func (storage *ProductStorageImpl) Delete(id uint) error {
	result := storage.db.Delete(&ProductModel{}, id)
	return result.Error
}
