package invoice

type CreateInvoiceDTO struct {
	CustomerName    string                    `json:"customerName" binding:"required" valid:"required~customerName is required"`
	CustomerPhoneNo string                    `json:"customerPhoneNo" binding:"required" valid:"required~customerPhoneNo is required"`
	BillingAddress  string                    `json:"billingAddress" binding:"required" valid:"required~billingAddress is required"`
	ShippingAddress string                    `json:"shippingAddress" binding:"required" valid:"required~shippingAddress is required"`
	InvoiceItems    []CreateInvoiceItemEntity `json:"orderItems" binding:"required"`
}

type CreateInvoiceItemEntity struct {
	ProductId int `json:"productId" binding:"required" valid:"required~productId is required,numeric~productId must be a valid interger value."`
	Quantity  int `json:"quantity" binding:"required" valid:"required~quantity is required,numeric~quantity must be a valid interger value."`
}

func (dto *CreateInvoiceItemEntity) ToModel() (*InvoiceItemModel, error) {
	return &InvoiceItemModel{
		ProductId: uint(dto.ProductId),
		Quantity:  uint(dto.Quantity),
	}, nil
}

type UpdateInvoiceDTO struct {
	Status string `json:"status" binding:"required" valid:"required~status is required,status~status must be one of the following: (Pending / Processed / Shipped / Delivered / Cancelled)"`
}
