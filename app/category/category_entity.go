package category

import (
	"time"

	"gorm.io/gorm"
)

type CategoryEntity struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	ID        uint           `json:"id"`
}

func CategoryEntityFromCategoryModel(model *CategoryModel) *CategoryEntity {
	return model.ToEntity()
}
