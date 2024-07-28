package middlewares

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
)

func ValidationMiddleware(payload interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(payload); err != nil {
			c.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse("Invalid body payload please provide all required fields."))
			c.Abort()
			return
		}
		_, err := govalidator.ValidateStruct(payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse(err.Error()))
			c.Abort()
		}
		fmt.Println("Going next")
		c.Next()
	}
}
