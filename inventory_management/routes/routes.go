package routes

import (
	"inventory_management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/produk", controllers.TambahProduk)
	r.GET("/produk", controllers.LihatSemuaProduk)
	r.GET("/produk/:id", controllers.LihatProdukByID)
	r.GET("/produk/kategori/:kategori", controllers.LihatProdukByKategori)
	r.PUT("/produk/:id", controllers.UpdateProduk)
	r.DELETE("/produk/:id", controllers.HapusProduk)

	r.GET("/inventaris/:id_produk", controllers.LihatStok)
	r.PATCH("/inventaris/:id_produk", controllers.UpdateStok)

	r.POST("/pesanan", controllers.BuatPesanan)
	r.GET("/pesanan/:id", controllers.LihatPesanan)

	return r
}
