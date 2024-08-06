package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-invoice-app/app/category"
	"github.com/kaungmyathan22/golang-invoice-app/app/common"
	"github.com/kaungmyathan22/golang-invoice-app/app/middlewares"
	"github.com/kaungmyathan22/golang-invoice-app/app/order"
	"github.com/kaungmyathan22/golang-invoice-app/app/product"
	"github.com/kaungmyathan22/golang-invoice-app/app/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IsNull(str string) bool {
	return len(str) == 0
}
func IsDecimal(str string) bool {
	if IsNull(str) {
		return false
	}
	_, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return true
	}
	_, err = strconv.ParseFloat(str, 64)
	return err == nil
}

func sixToEightDigitAlphanumericPasswordValidator(password string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{6,8}$`)
	return re.MatchString(password)
}

func main() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.TagMap["sixToEightDigitAlphanumericPasswordValidator"] = govalidator.Validator(sixToEightDigitAlphanumericPasswordValidator)
	govalidator.TagMap["isDecimal"] = govalidator.Validator(IsDecimal)

	dsn := "host=localhost user=admin password=admin dbname=invoice_app port=5433 sslmode=disable"
	// docker exec -it postgres psql -U admin -d postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("successfully connected to database.")
	db.AutoMigrate(&user.UserModel{})
	db.AutoMigrate(&user.PasswordResetTokenModel{})
	db.AutoMigrate(&category.CategoryModel{})
	db.AutoMigrate(&product.ProductModel{})
	db.AutoMigrate(&order.OrderModel{})
	db.AutoMigrate(&order.OrderItemModel{})

	r := gin.Default()
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, common.GetEnvelope(http.StatusNotFound, fmt.Sprintf("%s %s not found", ctx.Request.Method, ctx.Request.URL)))
	})
	v1Route := r.Group("/api/v1")
	/** user region */
	userStorage := user.NewUserStorage(db)
	tokenStorage := user.NewTokenStorage(db)
	userHandler := user.NewUserHandler(userStorage, tokenStorage)
	userRoutes := v1Route.Group("/auth")

	userRoutes.POST("/register", middlewares.ValidationMiddleware(&user.RegisterUserDTO{}), userHandler.CreateUserHandler)
	userRoutes.POST("/login", middlewares.ValidationMiddleware(&user.LoginUserDTO{}), userHandler.LoginHandler)

	userRoutes.GET("/users", userHandler.GetUsersHandler)
	userRoutes.GET("/me", middlewares.AuthMiddleware(userStorage), userHandler.MeHandler)

	userRoutes.PATCH("/change-password", middlewares.AuthMiddleware(userStorage), middlewares.ValidationMiddleware(&user.ChangePasswordDTO{}), userHandler.ChangePasswordHandler)
	userRoutes.PATCH("/profile", middlewares.AuthMiddleware(userStorage), middlewares.ValidationMiddleware(&user.UpdateUserDTO{}), userHandler.UpdateUserHandler)

	userRoutes.DELETE("/", middlewares.AuthMiddleware(userStorage), userHandler.DeleteUserHandler)

	userRoutes.POST("/forgot-password", middlewares.ValidationMiddleware(&user.ForgotPasswordDTO{}), userHandler.ForgotPasswordHandler)
	userRoutes.POST("/reset-password", middlewares.ValidationMiddleware(&user.ResetPasswordDTO{}), userHandler.ResetPasswordHandler)

	categoryStorage := category.NewCategoryStorage(db)
	categoryHandler := category.NewCategoryHandler(categoryStorage)
	categoryRoutes := v1Route.Group("/category")
	categoryRoutes.Use(middlewares.AuthMiddleware(userStorage))
	{
		categoryRoutes.POST("/", middlewares.ValidationMiddleware(&category.CreateCategoryDTO{}), categoryHandler.CreateCategoryHandler)
		categoryRoutes.GET("/", categoryHandler.GetCategoriesHandler)
		categoryRoutes.GET("/:id", categoryHandler.GetCategoryHandler)
		categoryRoutes.PATCH("/:id", middlewares.ValidationMiddleware(&category.UpdateCategoryDTO{}), categoryHandler.UpdateCategoryHandler)
		categoryRoutes.DELETE("/:id", categoryHandler.DeleteCategoryHandler)
	}

	/** Product  */
	productStorage := product.NewProductStorage(db)
	productHandler := product.NewProductHandler(productStorage)
	productRoutes := v1Route.Group("/products")
	productRoutes.Use(middlewares.AuthMiddleware(userStorage))
	{
		productRoutes.POST("/", middlewares.ValidationMiddleware(&product.CreateProductDTO{}), productHandler.CreateProductHandler)
		productRoutes.GET("/", productHandler.GetProductsHandler)
		productRoutes.GET("/:id", productHandler.GetProductHandler)
		productRoutes.PATCH("/:id", middlewares.ValidationMiddleware(&product.UpdateProductDTO{}), productHandler.UpdateProductHandler)
		productRoutes.DELETE("/:id", productHandler.DeleteProductHandler)
	}

	/** Order  */
	orderStorage := order.NewOrderStorage(db)
	orderHandler := order.NewOrderHandler(orderStorage, productStorage, db)
	orderRoutes := v1Route.Group("/orders")
	orderRoutes.Use(middlewares.AuthMiddleware(userStorage))
	{
		orderRoutes.POST("/", middlewares.ValidationMiddleware(&order.CreateOrderDTO{}), orderHandler.CreateOrderHandler)
		orderRoutes.GET("/", orderHandler.GetOrdersHandler)
		orderRoutes.GET("/:id", orderHandler.GetOrderHandler)
		// orderRoutes.PATCH("/:id", middlewares.ValidationMiddleware(&order.UpdateOrderDTO{}), orderHandler.UpdateOrderHandler)
		// orderRoutes.DELETE("/:id", orderHandler.DeleteOrderHandler)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	r.Run()
}
