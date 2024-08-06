package product

import (
	"errors"
	"fmt"
	"time"

	"github.com/kaungmyathan22/golang-invoice-app/app/category"
	"github.com/kaungmyathan22/golang-invoice-app/app/lib"
	"gorm.io/gorm"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product name already exists")
)

type ProductModel struct {
	ID         uint                   `gorm:"primaryKey"`
	Name       string                 `gorm:"column:name;;not null"`
	CategoryID *uint                  `gorm:"column:categoryId"`
	Category   category.CategoryModel `gorm:"constraint:OnDelete:SET NULL;"`
	Price      float64                `gorm:"type:decimal(10,2);column:price"`
	Slug       string                 `goorm:"column:slug;unique; not null"`
	CreatedAt  time.Time

	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductModel) TableName() string {
	return "product"
}

func (p *ProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	p.Slug = lib.GenerateSlug(p.Name)
	for {
		var count int64
		tx.Model(&ProductModel{}).Where("slug = ?", p.Slug).Count(&count)
		if count == 0 {
			break
		}
		p.Slug = lib.GenerateUniqueSlug(p.Slug)
	}
	fmt.Println(p.Slug)
	return nil
}

func (model *ProductModel) ToEntity() *ProductEntity {
	return &ProductEntity{
		ID:         model.ID,
		Name:       model.Name,
		Slug:       model.Slug,
		CategoryID: model.CategoryID,
		Price:      model.Price,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
		DeletedAt:  model.DeletedAt,
	}
}
