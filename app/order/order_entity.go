package order

import (
	"time"

	"github.com/kaungmyathan22/golang-invoice-app/app/product"
	"gorm.io/gorm"
)

type OrderEntity struct {
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
	OrderItems      *[]OrderItemEntity `json:"orderItems"`
	DeletedAt       gorm.DeletedAt     `json:"deletedAt"`
	OrderNo         string             `json:"orderNo"`
	OrderStatus     string             `json:"orderStatus"`
	CustomerName    string             `json:"customerName"`
	CustomerPhoneNo string             `json:"customerPhoneNo"`
	BillingAddress  string             `json:"billingAddress"`
	ShippingAddress string             `json:"shippingAddress"`
	ID              uint               `json:"id"`
	ShippingCosts   float64            `json:"shippingCosts"`
	SubTotal        float64            `json:"subTotal"`
	Total           float64            `json:"total"`
}

type OrderItemEntity struct {
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	Product   *product.ProductEntity `json:"product,omitempty"`
	DeletedAt gorm.DeletedAt         `json:"deletedAt"`
	ID        uint                   `json:"id"`
	OrderId   uint                   `json:"orderId"`
	ProductId uint                   `json:"productId"`
	Quantity  uint                   `json:"quantity"`
	Total     float64                `json:"total"`
}
