package order

import (
	"time"

	"github.com/kaungmyathan22/golang-invoice-app/app/product"
	"gorm.io/gorm"
)

type OrderEntity struct {
	ID              uint             `json:"id"`
	OrderNo         string           `json:"orderNo"`
	OrderStatus     string           `json:"orderStatus"`
	CustomerName    string           `json:"customerName"`
	CustomerPhoneNo string           `json:"customerPhoneNo"`
	BillingAddress  string           `json:"billingAddress"`
	ShippingAddress string           `json:"shippingAddress"`
	ShippingCosts   float64          `json:"shippingCosts"`
	OrderItems      *OrderItemEntity `json:"orderItems"`
	SubTotal        float64          `json:"subTotal"`
	Total           float64          `json:"total"`
	CreatedAt       time.Time        `json:"createdAt"`
	UpdatedAt       time.Time        `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt   `json:"deletedAt"`
}

func (entity *OrderEntity) ToModel() *OrderModel {
	return &OrderModel{
		ID:              entity.ID,
		OrderStatus:     OrderStatus(entity.OrderStatus),
		CustomerName:    entity.CustomerName,
		CustomerPhoneNo: entity.CustomerPhoneNo,
		BillingAddress:  entity.BillingAddress,
		ShippingAddress: entity.ShippingAddress,
		ShippingCosts:   entity.ShippingCosts,
		SubTotal:        entity.SubTotal,
		Total:           entity.Total,
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
		DeletedAt:       entity.DeletedAt,
	}
}

type OrderItemEntity struct {
	ID        uint                   `json:"id"`
	OrderId   uint                   `json:"orderId"`
	ProductId uint                   `json:"productId"`
	Product   *product.ProductEntity `json:"product"`
	Quantity  uint                   `json:"quantity"`
	Total     float64                `json:"total"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	DeletedAt gorm.DeletedAt         `json:"deletedAt"`
}

func (entity *OrderItemEntity) ToModel() *OrderItemModel {
	return &OrderItemModel{
		ID:        entity.ID,
		ProductId: entity.ProductId,
		OrderId:   entity.OrderId,
		Total:     entity.Total,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}
