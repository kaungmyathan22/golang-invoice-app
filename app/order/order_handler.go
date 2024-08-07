package order

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
	"github.com/kaungmyathan22/golang-invoice-app/app/product"
	"gorm.io/gorm"
)

type OrderHandler struct {
	orderStorage   OrderStorage
	productStorage product.ProductStorage
	DB             *gorm.DB
}

func NewOrderHandler(orderStorage OrderStorage, productStorage product.ProductStorage, db *gorm.DB) *OrderHandler {
	return &OrderHandler{orderStorage: orderStorage, productStorage: productStorage, DB: db}
}

func (handler *OrderHandler) GetOrdersHandler(ctx *gin.Context) {
	var pagination common.PaginationParamsRequest

	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, common.GetEnvelope(http.StatusBadRequest, err.Error()))
		return
	}
	pagination.SetDefaultPaginationValues()
	orders, err := handler.orderStorage.GetAll(pagination.Page, pagination.PageSize)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	var OrderEntities []OrderEntity
	for _, model := range orders {
		OrderEntities = append(OrderEntities, *model.ToEntity())
	}
	totalItems, err := handler.orderStorage.GetCount(nil)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, gin.H{
		"meta":   pagination.GetMeta(totalItems),
		"orders": OrderEntities,
	}))
}

func (handler *OrderHandler) CreateOrderHandler(ctx *gin.Context) {
	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*CreateOrderDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	order := payload.ToModel()
	err := handler.DB.Transaction(func(tx *gorm.DB) error {
		var orderItemModels []OrderItemModel
		for _, v := range payload.OrderItems {
			productModel, err := handler.productStorage.GetById(uint(v.ProductId))
			if err != nil {
				if errors.Is(err, product.ErrProductNotFound) {
					ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, err.Error()))
					return err
				}
				log.Println(err.Error())
				ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong"))
				return err
			}
			order.SubTotal += float64(float64(v.Quantity) * productModel.Price)
			orderItem := OrderItemModel{ProductId: productModel.ID, Quantity: uint(v.Quantity), Total: float64(float64(v.Quantity) * productModel.Price)}
			orderItemModels = append(orderItemModels, orderItem)
		}
		order.OrderStatus = string(OrderStatusPending)
		order.Total = order.ShippingCosts + order.SubTotal
		orderModel, err := handler.orderStorage.Create(*order)
		var responsePayload OrderEntity
		if err != nil {
			if strings.Contains(err.Error(), "(SQLSTATE 23503)") {
				ctx.JSON(http.StatusConflict, common.GetEnvelope(http.StatusConflict, "product with given id doesn't exists."))
				return err
			}
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong while creating order."))
			return err
		}
		responsePayload = *orderModel.ToEntity()
		var orderItemEntities []OrderItemEntity
		for k := range orderItemModels {
			orderItemModels[k].Order = *orderModel
			orderItemModels[k].OrderId = orderModel.ID
			orderItem, err := handler.orderStorage.CreateOrderItem(orderItemModels[k])
			if err != nil {
				log.Println(err.Error())
				return err
			}
			orderItemModels[k] = *orderItem
			orderItemEntities = append(orderItemEntities, *orderItem.ToEntity())
		}
		responsePayload.OrderItems = &orderItemEntities
		ctx.JSON(http.StatusCreated, common.GetEnvelope(http.StatusCreated, responsePayload))
		return nil
	})
	if err != nil {
		log.Println(err.Error())
	}
}

func (handler *OrderHandler) DeleteOrderHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	orderId := uint(id)

	order, err := handler.orderStorage.GetById(orderId)
	if err != nil {
		if errors.Is(err, ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Order with given id %s not found", idStr)))
			return
		}
	}
	if err := handler.orderStorage.Delete((order.ID)); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong."))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully deleted Order."))
}

func (handler *OrderHandler) GetOrderHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	orderId := uint(id)

	order, err := handler.orderStorage.GetById(orderId)
	if err != nil {
		if errors.Is(err, ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Order with given id %s not found", idStr)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong whlie getting order"))
		return
	}
	orderItems, err := handler.orderStorage.GetOrderItems(order.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong whlie getting order items"))
		return
	}
	var orderItemsEntities []OrderItemEntity
	for _, model := range orderItems {
		orderItemsEntities = append(orderItemsEntities, *model.ToEntity())
	}
	orderEntity := order.ToEntity()
	orderEntity.OrderItems = &orderItemsEntities
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, orderEntity))
}

func (handler *OrderHandler) UpdateOrderHandler(ctx *gin.Context) {
	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*UpdateOrderDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	orderId := uint(id)

	order, err := handler.orderStorage.GetById(orderId)
	if err != nil {
		if errors.Is(err, ErrOrderNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Order with given id %s not found", idStr)))
			return
		}
	}
	order.OrderStatus = payload.Status
	err = handler.orderStorage.Update(*order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong while changing Order status"))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully updated order status"))
}
