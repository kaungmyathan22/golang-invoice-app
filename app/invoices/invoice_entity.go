package invoice

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceItemEntity struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	ID        uint           `json:"id"`
}
