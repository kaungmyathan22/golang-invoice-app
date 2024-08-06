package product

import (
	"time"

	"gorm.io/gorm"
)

type ProductEntity struct {
	ID         uint           `json:"id"`
	Name       string         `json:"name"`
	Slug       string         `json:"slug"`
	CategoryID *uint          `json:"categoryId"`
	Price      float64        `json:"price"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
}

func (entity *ProductEntity) ToModel(model *ProductModel) *ProductModel {
	return &ProductModel{
		Name:       entity.Name,
		CategoryID: entity.CategoryID,
		Price:      entity.Price,
	}
}
