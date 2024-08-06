package order

type CreateOrderDTO struct {
	Name       string  `json:"name" binding:"required" valid:"required~Order name is required"`
	Price      float64 `json:"price" binding:"required" valid:"required~Order price is required,isDecimal~Price must be greater than zero"`
	CategoryID uint    `json:"categoryId" binding:"required" valid:"required~Category ID is required,isDecimal~categoryId must be valid number"`
}

func (dto *CreateOrderDTO) ToModel() (*OrderModel, error) {
	return &OrderModel{}, nil
}

type UpdateOrderDTO struct {
	CreateOrderDTO
}
