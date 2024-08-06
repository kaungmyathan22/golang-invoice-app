package product

import (
	"time"

	"gorm.io/gorm"
)

type ProductEntity struct {
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	CategoryID *uint          `json:"categoryId"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
	Name       string         `json:"name"`
	Slug       string         `json:"slug"`
	ID         uint           `json:"id"`
	Price      float64        `json:"price"`
}

func (entity *ProductEntity) ToModel(model *ProductModel) *ProductModel {
	return &ProductModel{
		Name:       entity.Name,
		CategoryID: entity.CategoryID,
		Price:      entity.Price,
	}
}
