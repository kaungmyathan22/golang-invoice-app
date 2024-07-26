package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	user_models "github.com/kaungmyathan22/golang-invoice-app/app/user/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// config.InitApp()
	dsn := "host=localhost user=admin password=admin dbname=invoice_app port=5433 sslmode=disable"
	// docker exec -it postgres psql -U admin -d postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("successfully connected to database.")
	db.AutoMigrate(&user_models.UserModel{})

	r := gin.Default()
	v1Route := r.Group("/api/v1")
	userRoutes := v1Route.Group("/users")
	userRoutes.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola"})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
