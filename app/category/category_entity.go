package category

import (
	"time"

	"gorm.io/gorm"
)

type CategoryEntity struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func CategoryEntityFromCategoryModel(model *CategoryModel) *CategoryEntity {
	return &CategoryEntity{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
		Name:      model.Name,
	}
}
