package invoice

import (
	"errors"
	"time"

	"github.com/kaungmyathan22/golang-invoice-app/app/order"
	"github.com/kaungmyathan22/golang-invoice-app/app/product"
	"gorm.io/gorm"
)

var (
	ErrInvoiceNotFound      = errors.New("invoice not found")
	ErrInvoiceAlreadyExists = errors.New("invoice name already exists")
)

type InvoiceItemModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt       `gorm:"index"`
	Product   product.ProductModel `gorm:"constraint:OnDelete:RESTRICT;"`
	Order     order.OrderModel     `gorm:"constraint:OnDelete:CASCADE;"`
	ID        uint                 `gorm:"primaryKey"`
	InvoiceId uint                 `gorm:"column:orderId;not null"`
	ProductId uint                 `gorm:"column:productId;not null"`
	Quantity  uint                 `gorm:"column:quantity;not null"`
	Total     float64              `gorm:"column:total;not null"`
}

func (InvoiceItemModel) TableName() string {
	return "invoices"
}
