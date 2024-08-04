package category

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrCategoryNotFound      = errors.New("Category not found")
	ErrCategoryAlreadyExists = errors.New("Category name already exists")
)

type CategoryModel struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"column:username;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (CategoryModel) TableName() string {
	return "category"
}

func (model *CategoryModel) ToEntity() *CategoryEntity {
	return &CategoryEntity{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}
