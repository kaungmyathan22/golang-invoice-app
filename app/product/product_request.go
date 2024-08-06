package product

type CreateProductDTO struct {
	Name       string  `json:"name" binding:"required" valid:"required~Product name is required"`
	Price      float64 `json:"price" binding:"required" valid:"required~Product price is required,isDecimal~Price must be greater than zero"`
	CategoryID uint    `json:"categoryId" binding:"required" valid:"required~Category ID is required,isDecimal~categoryId must be valid number"`
}

func (dto *CreateProductDTO) ToModel() (*ProductModel, error) {
	return &ProductModel{Name: dto.Name, Price: dto.Price, CategoryID: &dto.CategoryID}, nil
}

type UpdateProductDTO struct {
	CreateProductDTO
}
