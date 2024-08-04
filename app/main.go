package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
	"github.com/kaungmyathan22/golang-invoice-app/app/middlewares"
	"github.com/kaungmyathan22/golang-invoice-app/app/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func sixToEightDigitAlphanumericPasswordValidator(password string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{6,8}$`)
	return re.MatchString(password)
}

func main() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.TagMap["sixToEightDigitAlphanumericPasswordValidator"] = govalidator.Validator(sixToEightDigitAlphanumericPasswordValidator)

	dsn := "host=localhost user=admin password=admin dbname=invoice_app port=5433 sslmode=disable"
	// docker exec -it postgres psql -U admin -d postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("successfully connected to database.")
	db.AutoMigrate(&user.UserModel{})
	db.AutoMigrate(&user.PasswordResetTokenModel{})

	r := gin.Default()
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("%s %s not found", ctx.Request.Method, ctx.Request.URL)))
	})
	v1Route := r.Group("/api/v1")
	/** user region */
	userStorage := user.NewUserStorage(db)
	userHandler := user.NewUserHandler(userStorage)
	userRoutes := v1Route.Group("/auth")

	userRoutes.POST("/register", middlewares.ValidationMiddleware(&user.RegisterUserDTO{}), userHandler.CreateUserHandler)
	userRoutes.POST("/login", middlewares.ValidationMiddleware(&user.LoginUserDTO{}), userHandler.LoginHandler)

	userRoutes.GET("/users", userHandler.GetUsersHandler)
	userRoutes.GET("/me", middlewares.AuthMiddleware(userStorage), userHandler.MeHandler)

	userRoutes.PATCH("/change-password", middlewares.AuthMiddleware(userStorage), middlewares.ValidationMiddleware(&user.ChangePasswordDTO{}), userHandler.ChangePasswordHandler)
	userRoutes.PATCH("/profile", middlewares.AuthMiddleware(userStorage), middlewares.ValidationMiddleware(&user.UpdateUserDTO{}), userHandler.UpdateUserHandler)

	userRoutes.DELETE("/", middlewares.AuthMiddleware(userStorage), userHandler.DeleteUserHandler)

	userRoutes.POST("/forgot-password", middlewares.ValidationMiddleware(&user.ForgotPasswordDTO{}), userHandler.ForgotPasswordHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	r.Run()
}
