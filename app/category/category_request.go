package category

type CreateCategoryDTO struct {
	Name string `json:"name" binding:"required" valid:"required~Category name is required"`
}

func (dto *CreateCategoryDTO) ToModel() (*CategoryModel, error) {
	return &CategoryModel{Name: dto.Name}, nil
}

type UpdateCategoryDTO struct {
	Name string `json:"name" binding:"required" valid:"required~Category name is required"`
}
