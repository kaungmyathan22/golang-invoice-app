package user_handlers

import (
	"github.com/gin-gonic/gin"
	user_models "github.com/kaungmyathan22/golang-invoice-app/app/user/models"
)

type UserHandler struct {
	Storage user_models.UserStorage
}

func NewUserHandler(db user_models.UserStorage) *UserHandler {
	return &UserHandler{Storage: db}
}

func (handler *UserHandler) GetUsersHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetUsersHandler"})
}
func (handler *UserHandler) GetUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetUserHandler"})
}
func (handler *UserHandler) CreateUserHandler(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Register Successful"})
}
func (handler *UserHandler) UpdateUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "UpdateUserHandler"})
}
func (handler *UserHandler) DeleteUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "DeleteUserHandler"})
}
