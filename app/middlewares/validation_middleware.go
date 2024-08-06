package middlewares

import (
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
)

func ValidationMiddleware(payload interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(payload); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, common.GetEnvelope(http.StatusUnprocessableEntity, "Invalid body payload please provide all required fields."))
			c.Abort()
			return
		}
		isValid, err := govalidator.ValidateStruct(payload)
		if err != nil || !isValid {
			log.Println(err.Error())
			c.JSON(http.StatusUnprocessableEntity, common.GetEnvelope(http.StatusUnprocessableEntity, err.Error()))
			c.Abort()
		}
		c.Set("payload", payload)
		c.Next()
	}
}
