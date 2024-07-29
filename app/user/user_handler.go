package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
)

type UserHandler struct {
	Storage UserStorage
}

func NewUserHandler(db UserStorage) *UserHandler {
	return &UserHandler{Storage: db}
}

func (handler *UserHandler) GetUsersHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetUsersHandler"})
}
func (handler *UserHandler) GetUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetUserHandler"})
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
	ctx.JSON(200, gin.H{"message": "UpdateUserHandler"})
}
func (handler *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "DeleteUserHandler"})
}
