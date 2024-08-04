package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
	"github.com/kaungmyathan22/golang-invoice-app/app/lib"
)

type UserHandler struct {
	Storage UserStorage
}

func NewUserHandler(db UserStorage) *UserHandler {
	return &UserHandler{Storage: db}
}

func (handler *UserHandler) GetUsersHandler(ctx *gin.Context) {
	var pagination common.PaginationParamsRequest

	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, common.GetEnvelope(http.StatusBadRequest, err.Error()))
		return
	}
	pagination.SetDefaultPaginationValues()
	users, err := handler.Storage.GetAll(pagination.Page, pagination.PageSize)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	var userEntities []UserEntity
	for _, model := range users {
		userEntities = append(userEntities, *UserEntityFromUserModel(&model))
	}
	totalItems, err := handler.Storage.GetCount(nil)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, nil))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, gin.H{
		"meta":  pagination.GetMeta(totalItems),
		"users": userEntities,
	}))
}

func (handler *UserHandler) LoginHandler(ctx *gin.Context) {
	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*LoginUserDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	user, err := handler.Storage.GetByUsername(payload.Username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, common.GetStatusBadRequestResponse("Invalid username / password"))
			return
		}
	}
	if !lib.CheckPasswordHash(payload.Password, user.Password) {
		ctx.JSON(http.StatusNotFound, common.GetStatusBadRequestResponse("Invalid username / password"))
		return
	}
	token, err := lib.GenerateToken(user.ID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, common.GetInternalServerErrorResponse("Something went wrong."))
		return
	}
	ctx.JSON(200, common.GetSuccessResponse(UserLoginResponse{User: *UserEntityFromUserModel(user), Token: token}))
}

func (handler *UserHandler) MeHandler(ctx *gin.Context) {
	userModel, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}
	ctx.JSON(http.StatusOK, common.GetSuccessResponse(UserEntityFromUserModel(userModel.(*UserModel))))
}

func (handler *UserHandler) CreateUserHandler(c *gin.Context) {
	rawPayload, exists := c.Get("payload")
	if !exists {
		c.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*RegisterUserDTO)
	if !ok {
		c.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	user, err := payload.ToModel()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("can't hash password as expected"))
		return
	}

	user, err = handler.Storage.Create(*user)
	if err != nil {
		if common.IsUniqueKeyViolation(err) {
			c.JSON(http.StatusConflict, common.GetStatusConflictResponse(ErrUsernameAlreadyExists.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusAccepted, common.GetSuccessResponse(FromModel(user)))
}

func (handler *UserHandler) UpdateUserHandler(ctx *gin.Context) {
	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*UpdateUserDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	rawUser, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}
	userModel, ok := rawUser.(*UserModel)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}
	userModel.Username = payload.Username
	err := handler.Storage.Update(*userModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong while changing username"))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully changed username"))
}

func (handler *UserHandler) ChangePasswordHandler(ctx *gin.Context) {
	rawPayload, exists := ctx.Get("payload")
	if !exists {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("payload do not exists"))
	}
	payload, ok := rawPayload.(*ChangePasswordDTO)
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("invalid payload type"))
		return
	}
	if payload.OldPassword == payload.NewPassword {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("newPassword can't be the same with oldPassword"))
		return
	}
	rawUser, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}
	userModel, ok := rawUser.(*UserModel)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}
	if !lib.CheckPasswordHash(payload.OldPassword, userModel.Password) {
		ctx.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("incorrect old password"))
		return
	}
	err := payload.HashPassword()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong while hashing password"))
		return
	}
	userModel.Password = payload.NewPassword
	err = handler.Storage.Update(*userModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong while changing password"))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, "Successfully changed password"))
}

func (handler *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	rawUser, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}
	userModel, ok := rawUser.(*UserModel)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, common.GetInternalServerErrorResponse("something went wrong"))
		return
	}
	if err := handler.Storage.Delete((userModel.ID)); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, gin.H{"message": "Something went wrong."}))
		return
	}
	ctx.JSON(http.StatusOK, common.GetEnvelope(http.StatusOK, gin.H{"message": "Successfully deleted user."}))
}
