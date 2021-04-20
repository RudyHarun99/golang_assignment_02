package config

import (
	"fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
	"GO14/assignment02/models"
)

var (
	Db    *gorm.DB
	errDB error
)

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=orders_by port=5432 sslmode=disable"
	Db, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		fmt.Println("error:", errDB)
	}

	fmt.Println("db connection established")
	Db.AutoMigrate(&models.Orders{})
	Db.AutoMigrate(&models.Item{})
}