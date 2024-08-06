package order

import (
	"errors"
	"time"

	"github.com/kaungmyathan22/golang-invoice-app/app/lib"
	"github.com/kaungmyathan22/golang-invoice-app/app/product"
	"gorm.io/gorm"
)

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order name already exists")
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusProcessed OrderStatus = "Processed"
	OrderStatusShipped   OrderStatus = "Shipped"
	OrderStatusDelivered OrderStatus = "Delivered"
	OrderStatusCancelled OrderStatus = "Cancelled"
)

// func (os *OrderStatus) Scan(value interface{}) error {
// 	*os = OrderStatus(value.([]byte))
// 	fmt.Println(value)
// 	return nil
// }

// func (os OrderStatus) Value() (driver.Value, error) {
// 	return string(os), nil
// }

// func (os OrderStatus) IsValid() bool {
// 	switch os {
// 	case OrderStatusPending, OrderStatusProcessed, OrderStatusShipped, OrderStatusDelivered, OrderStatusCancelled:
// 		return true
// 	}
// 	return false
// }

// func (os OrderStatus) String() string {
// 	return string(os)
// }

/**
 * CREATE TYPE order_status AS ENUM ('Pending', 'Processed', 'Shipped', 'Delivered', 'Cancelled');
 */
type OrderModel struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	OrderNo         string         `gorm:"column:order_no;not null;unique"`
	OrderStatus     string         `gorm:"type:order_status;not null; default:'Pending'"`
	CustomerName    string         `gorm:"column:customer_name;not null"`
	CustomerPhoneNo string         `gorm:"column:customer_phoneNo;not null"`
	BillingAddress  string         `gorm:"column:billing_address;not null"`
	ShippingAddress string         `gorm:"column:shipping_address;not null"`
	ID              uint           `gorm:"primaryKey"`
	ShippingCosts   float64        `gorm:"column:shipping_costs;not null"`
	SubTotal        float64        `gorm:"column:sub_total;not null"`
	Total           float64        `gorm:"column:total;not null"`
}

func (OrderModel) TableName() string {
	return "order"
}

func (order *OrderModel) BeforeCreate(tx *gorm.DB) (err error) {
	order.OrderNo, _ = lib.GenerateOrderNumber()
	for {
		var count int64
		tx.Model(&OrderModel{}).Where("order_no = ?", order.OrderNo).Count(&count)
		if count == 0 {
			break
		}
		order.OrderNo, _ = lib.GenerateOrderNumber()
	}
	return nil
}

type OrderItemModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt       `gorm:"index"`
	Product   product.ProductModel `gorm:"constraint:OnDelete:RESTRICT;"`
	Order     OrderModel           `gorm:"constraint:OnDelete:CASCADE;"`
	ID        uint                 `gorm:"primaryKey"`
	OrderId   uint                 `gorm:"column:orderId;not null"`
	ProductId uint                 `gorm:"column:productId;not null"`
	Quantity  uint                 `gorm:"column:quantity;not null"`
	Total     float64              `gorm:"column:total;not null"`
}

func (OrderItemModel) TableName() string {
	return "order_item"
}
