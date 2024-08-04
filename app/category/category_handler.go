package category

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
)

type CategoryHandler struct {
	categoryStorage CategoryStorage
}

func NewCategoryHandler(categoryStorage CategoryStorage) *CategoryHandler {
	return &CategoryHandler{categoryStorage: categoryStorage}
}

func (handler *CategoryHandler) GetCategoriesHandler(ctx *gin.Context) {
	var pagination common.PaginationParamsRequest

	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, common.GetEnvelope(http.StatusBadRequest, err.Error()))
		return
	}
	pagination.SetDefaultPaginationValues()
	categories, err := handler.categoryStorage.GetAll(pagination.Page, pagination.PageSize)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	var categoryEntities []CategoryEntity
	for _, model := range categories {
		categoryEntities = append(categoryEntities, *CategoryEntityFromCategoryModel(&model))
	}
	totalItems, err := handler.categoryStorage.GetCount(nil)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, gin.H{
		"meta":       pagination.GetMeta(totalItems),
		"categories": categoryEntities,
	}))
}

func (handler *CategoryHandler) CreateCategoryHandler(ctx *gin.Context) {
	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*CreateCategoryDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	category, err := payload.ToModel()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("can't hash password as expected"))
		return
	}

	category, err = handler.categoryStorage.Create(*category)
	if err != nil {
		if common.IsUniqueKeyViolation(err) {
			ctx.JSON(http.StatusConflict, common.GetStatusConflictResponse(ErrCategoryAlreadyExists.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusAccepted, common.GetSuccessResponse((category.ToEntity())))
}

func (handler *CategoryHandler) UpdateCategoryHandler(ctx *gin.Context) {

	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*UpdateCategoryDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	categoryId := uint(id)

	category, err := handler.categoryStorage.GetById(categoryId)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("category with given id %s not found", idStr)))
			return
		}
	}
	category.Name = payload.Name
	err = handler.categoryStorage.Update(*category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong while changing category name"))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully changed category name"))
}

func (handler *CategoryHandler) DeleteCategoryHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	categoryId := uint(id)

	category, err := handler.categoryStorage.GetById(categoryId)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("category with given id %s not found", idStr)))
			return
		}
	}
	if err := handler.categoryStorage.Delete((category.ID)); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "Something went wrong."))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, gin.H{"message": "Successfully deleted category."}))
}
func (handler *CategoryHandler) GetCategoryHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	categoryId := uint(id)

	category, err := handler.categoryStorage.GetById(categoryId)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("category with given id %s not found", idStr)))
			return
		}
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, category))
}
