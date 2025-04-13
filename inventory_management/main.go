package main

import (
	"inventory_management/controllers"
	"inventory_management/models"
	"inventory_management/routes"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDatabase() {
	dsn := "inventory_user:invpass123@tcp(127.0.0.1:3306)/inventory_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	database.AutoMigrate(&models.Produk{}, &models.Inventaris{}, &models.Pesanan{})
	// DB = database
	// log.Println("Database terkoneksi dan dimigrasi!")
	controllers.DB = database
}

func main() {
	initDatabase()

	r := routes.SetupRouter()
	r.Run(":8080")
}
