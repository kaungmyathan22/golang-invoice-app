package order

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
)

type OrderHandler struct {
	orderStorage OrderStorage
}

func NewOrderHandler(OrderStorage OrderStorage) *OrderHandler {
	return &OrderHandler{orderStorage: OrderStorage}
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
	// 	rawPayload, exists := ctx.Get("payload")
	// 	if !exists {
	// 		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	// 	}
	// 	payload, ok := rawPayload.(*CreateOrderDTO)
	// 	if !ok {
	// 		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
	// 		return
	// 	}
	// 	log.Println(payload.Price, payload.CategoryID)
	// 	order, err := payload.ToModel()
	// 	if err != nil {

	// 		log.Println(err.Error())
	// 		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
	// 		return
	// 	}

	// 	order, err = handler.OrderStorage.Create(*order)
	// 	if err != nil {
	// 		if common.IsUniqueKeyViolation(err) {
	// 			ctx.JSON(http.StatusConflict, common.GetStatusConflictResponse(ErrOrderAlreadyExists.Error()))
	// 			return
	// 		} else if strings.Contains(err.Error(), "(SQLSTATE 23503)") {
	// 			ctx.JSON(http.StatusConflict, common.GetEnvelope(http.StatusConflict, "category with given id doesn't exists."))
	// 			return
	// 		}
	// 		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse(err.Error()))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusCreated, common.GetEnvelope(http.StatusCreated, (order.ToEntity())))
	// }

	// func (handler *OrderHandler) UpdateOrderHandler(ctx *gin.Context) {

	// 	rawPayload, exists := ctx.Get("payload")
	// 	if !exists {
	// 		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	// 	}
	// 	payload, ok := rawPayload.(*UpdateOrderDTO)
	// 	if !ok {
	// 		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
	// 		return
	// 	}
	// 	idStr := ctx.Param("id")
	// 	id, err := strconv.ParseInt(idStr, 10, 32)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
	// 		return
	// 	}
	// 	orderId := uint(id)

	// 	order, err := handler.OrderStorage.GetById(orderId)
	// 	if err != nil {
	// 		if errors.Is(err, ErrOrderNotFound) {
	// 			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Order with given id %s not found", idStr)))
	// 			return
	// 		}
	// 	}
	// 	order.Name = payload.Name
	// 	order.CategoryID = &payload.CategoryID
	// 	err = handler.OrderStorage.Update(*order)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong while changing Order name"))
	// 		return
	// 	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully changed Order name"))
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
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, order.ToEntity()))
}
