package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	v1Route := r.Group("/api/v1")
	/** user region */
	userStorage := user.NewUserStorage(db)
	userHandler := user.NewUserHandler(userStorage)
	userRoutes := v1Route.Group("/auth")
	userRoutes.POST("/register", middlewares.ValidationMiddleware(&user.RegisterUserDTO{}), userHandler.CreateUserHandler)
	userRoutes.POST("/login", middlewares.ValidationMiddleware(&user.LoginUserDTO{}), userHandler.LoginHandler)
	userRoutes.GET("/me", middlewares.AuthMiddleware(userStorage), userHandler.MeHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	r.Run()
}
