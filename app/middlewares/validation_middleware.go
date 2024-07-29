package middlewares

import (
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
		isValid, err := govalidator.ValidateStruct(payload)
		if err != nil || !isValid {
			c.JSON(http.StatusBadRequest, common.GetStatusBadRequestResponse(err.Error()))
			c.Abort()
		}
		c.Set("payload", payload)
		c.Next()
	}
}
