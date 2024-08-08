package invoice

type CreateInvoiceDTO struct {
	CustomerName    string                    `json:"customerName" binding:"required" valid:"required~customerName is required"`
	CustomerPhoneNo string                    `json:"customerPhoneNo" binding:"required" valid:"required~customerPhoneNo is required"`
	BillingAddress  string                    `json:"billingAddress" binding:"required" valid:"required~billingAddress is required"`
	ShippingAddress string                    `json:"shippingAddress" binding:"required" valid:"required~shippingAddress is required"`
	InvoiceItems    []CreateInvoiceItemEntity `json:"orderItems" binding:"required"`
}

type CreateInvoiceItemEntity struct {
	OrderId int `json:"productId" binding:"required" valid:"required~productId is required,numeric~productId must be a valid interger value."`
}
