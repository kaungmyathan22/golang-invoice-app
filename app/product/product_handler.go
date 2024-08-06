package product

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
)

type ProductHandler struct {
	ProductStorage ProductStorage
}

func NewProductHandler(ProductStorage ProductStorage) *ProductHandler {
	return &ProductHandler{ProductStorage: ProductStorage}
}

func (handler *ProductHandler) GetProductsHandler(ctx *gin.Context) {
	var pagination common.PaginationParamsRequest

	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, common.GetEnvelope(http.StatusBadRequest, err.Error()))
		return
	}
	pagination.SetDefaultPaginationValues()
	products, err := handler.ProductStorage.GetAll(pagination.Page, pagination.PageSize)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	var ProductEntities []ProductEntity
	for _, model := range products {
		ProductEntities = append(ProductEntities, *model.ToEntity())
	}
	totalItems, err := handler.ProductStorage.GetCount(nil)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, gin.H{
		"meta":     pagination.GetMeta(totalItems),
		"products": ProductEntities,
	}))
}

func (handler *ProductHandler) CreateProductHandler(ctx *gin.Context) {
	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*CreateProductDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	log.Println(payload.Price, payload.CategoryID)
	product, err := payload.ToModel()
	if err != nil {

		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}

	product, err = handler.ProductStorage.Create(*product)
	if err != nil {
		if common.IsUniqueKeyViolation(err) {
			ctx.JSON(http.StatusConflict, common.GetStatusConflictResponse(ErrProductAlreadyExists.Error()))
			return
		} else if strings.Contains(err.Error(), "(SQLSTATE 23503)") {
			ctx.JSON(http.StatusConflict, common.GetEnvelope(http.StatusConflict, "category with given id doesn't exists."))
			return
		}
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, common.GetEnvelope(http.StatusCreated, (product.ToEntity())))
}

func (handler *ProductHandler) UpdateProductHandler(ctx *gin.Context) {

	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*UpdateProductDTO)
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
	productId := uint(id)

	product, err := handler.ProductStorage.GetById(productId)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Product with given id %s not found", idStr)))
			return
		}
	}
	product.Name = payload.Name
	product.CategoryID = &payload.CategoryID
	err = handler.ProductStorage.Update(*product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong while changing Product name"))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully changed Product name"))
}

func (handler *ProductHandler) DeleteProductHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	productId := uint(id)

	product, err := handler.ProductStorage.GetById(productId)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Product with given id %s not found", idStr)))
			return
		}
	}
	if err := handler.ProductStorage.Delete((product.ID)); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong."))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully deleted Product."))
}

func (handler *ProductHandler) GetProductHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	productId := uint(id)

	product, err := handler.ProductStorage.GetById(productId)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("Product with given id %s not found", idStr)))
			return
		}
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, product.ToEntity()))
}
