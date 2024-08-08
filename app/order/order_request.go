package order

type CreateOrderDTO struct {
	CustomerName    string                  `json:"customerName" binding:"required" valid:"required~customerName is required"`
	CustomerPhoneNo string                  `json:"customerPhoneNo" binding:"required" valid:"required~customerPhoneNo is required"`
	CustomerEmail   string                  `json:"customerEmail" binding:"required" valid:"required~customerEmail is required,email~Invalid email"`
	BillingAddress  string                  `json:"billingAddress" binding:"required" valid:"required~billingAddress is required"`
	ShippingAddress string                  `json:"shippingAddress" binding:"required" valid:"required~shippingAddress is required"`
	OrderItems      []CreateOrderItemEntity `json:"orderItems" binding:"required"`
}

type CreateOrderItemEntity struct {
	ProductId int `json:"productId" binding:"required" valid:"required~productId is required,numeric~productId must be a valid interger value."`
	Quantity  int `json:"quantity" binding:"required" valid:"required~quantity is required,numeric~quantity must be a valid interger value."`
}

func (dto *CreateOrderItemEntity) ToModel() (*OrderItemModel, error) {
	return &OrderItemModel{
		ProductId: uint(dto.ProductId),
		Quantity:  uint(dto.Quantity),
	}, nil
}

func (dto *CreateOrderDTO) ToModel() *OrderModel {
	return &OrderModel{
		CustomerName:    dto.CustomerName,
		CustomerEmail:   dto.CustomerEmail,
		CustomerPhoneNo: dto.CustomerPhoneNo,
		BillingAddress:  dto.BillingAddress,
		ShippingAddress: dto.ShippingAddress,
	}
}

type UpdateOrderDTO struct {
	Status string `json:"status" binding:"required" valid:"required~status is required,status~status must be one of the following: (Pending / Processed / Shipped / Delivered / Cancelled)"`
}
