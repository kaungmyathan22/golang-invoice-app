package category

import (
	"time"

	"gorm.io/gorm"
)

type CategoryEntity struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func CategoryEntityFromCategoryModel(model *CategoryModel) *CategoryEntity {
	return model.ToEntity()
}
