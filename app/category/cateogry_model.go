package category

import (
	"errors"
	"time"

	"github.com/kaungmyathan22/golang-invoice-app/app/lib"
	"gorm.io/gorm"
)

var (
	ErrCategoryNotFound      = errors.New("Category not found")
	ErrCategoryAlreadyExists = errors.New("Category name already exists")
)

type CategoryModel struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"column:name;not null"`
	Slug      string `gorm:"column:slug;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (CategoryModel) TableName() string {
	return "category"
}

func (p *CategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = lib.GenerateSlug(p.Name)
	for {
		var count int64
		tx.Model(&CategoryModel{}).Where("slug = ?", p.Slug).Count(&count)
		if count == 0 {
			break
		}
		p.Slug = lib.GenerateUniqueSlug(p.Slug)
	}
	return nil
}

func (model *CategoryModel) ToEntity() *CategoryEntity {
	return &CategoryEntity{
		ID:        model.ID,
		Name:      model.Name,
		Slug:      model.Slug,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}
