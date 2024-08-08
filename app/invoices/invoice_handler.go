package invoice

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
	"github.com/kaungmyathan22/golang-invoice-app/app/order"
	"gorm.io/gorm"
)

type InvoiceHandler struct {
	orderStorage order.OrderStorage
	DB           *gorm.DB
}

func NewInvoiceHandler(orderStorage order.OrderStorage) *InvoiceHandler {
	return &InvoiceHandler{orderStorage: orderStorage}
}

func (handler *InvoiceHandler) SendInvoicesHandler(ctx *gin.Context) {
	idStr := ctx.Param("orderId")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	orderId := uint(id)

	orderModel, err := handler.orderStorage.GetById(orderId)
	if err != nil {
		if errors.Is(err, order.ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Order with given id %s not found", idStr)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong whlie getting order"))
		return
	}
	orderItems, err := handler.orderStorage.GetOrderItems(orderModel.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong whlie getting order items"))
		return
	}
	var orderItemsEntities []order.OrderItemEntity
	for _, model := range orderItems {
		orderItemsEntities = append(orderItemsEntities, *model.ToEntity())
	}
	orderEntity := orderModel.ToEntity()
	orderEntity.OrderItems = &orderItemsEntities
	common.SendEmailHandler(common.EmailData{To: orderModel.CustomerEmail, Subject: "Order invoice", Template: common.INVOICE_EMAIL_TEMPLATE, Data: orderEntity})
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, orderEntity))
}
