package config

import (
	"log"

	"github.com/kaungmyathan22/golang-invoice-app/app/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=admin dbname=invoice_app port=5433 sslmode=disable"
	// docker exec -it postgres psql -U admin -d postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("successfully connected to database.")
	db.AutoMigrate(&user.UserModel{})

	return db, nil
}
