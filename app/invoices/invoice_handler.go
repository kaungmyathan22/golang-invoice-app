package invoice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
	"github.com/kaungmyathan22/golang-invoice-app/app/product"
	"gorm.io/gorm"
)

type InvoiceHandler struct {
	orderStorage   InvoiceStorage
	productStorage product.ProductStorage
	DB             *gorm.DB
}

func NewInvoiceHandler(orderStorage InvoiceStorage, productStorage product.ProductStorage, db *gorm.DB) *InvoiceHandler {
	return &InvoiceHandler{orderStorage: orderStorage, productStorage: productStorage, DB: db}
}

func (handler *InvoiceHandler) GetInvoicesHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, gin.H{
		"meta":     "meta",
		"invoices": "invoices",
	}))
}
