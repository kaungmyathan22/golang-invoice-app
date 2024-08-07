package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
	"github.com/kaungmyathan22/golang-invoice-app/app/lib"
	"github.com/kaungmyathan22/golang-invoice-app/app/user"
)

func AuthMiddleware(userStorage user.UserStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		log.Println(authHeader)
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, common.GetEnvelope(http.StatusUnauthorized, "Authorization header is missing"))
			ctx.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, common.GetEnvelope(http.StatusUnauthorized, "Authorization token is missing"))
			ctx.Abort()
			return
		}

		claims, err := lib.VerifyToken(tokenString)

		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusUnauthorized, common.GetEnvelope(http.StatusUnauthorized, "Invalid token"))
			ctx.Abort()
			return
		}

		user_model, err := userStorage.GetById(claims.UserID)
		if err != nil {
			if errors.Is(err, user.ErrUserNotFound) {
				ctx.JSON(http.StatusUnauthorized, common.GetEnvelope(http.StatusUnauthorized, "user with given id not found."))
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusInternalServerError, common.GetEnvelope(http.StatusInternalServerError, "something went wrong"))
			ctx.Abort()
			return
		}
		fmt.Println("Going next....")
		ctx.Set("user", user_model)
		ctx.Next()
	}
}
