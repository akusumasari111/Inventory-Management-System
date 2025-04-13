package controllers

import (
	"inventory_management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuatPesanan(c *gin.Context) {
	var pesanan models.Pesanan
	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var produk models.Produk
	if err := DB.First(&produk, pesanan.IDProduk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	DB.Create(&pesanan)
	c.JSON(http.StatusOK, pesanan)
}

func LihatPesanan(c *gin.Context) {
	var pesanan models.Pesanan
	id := c.Param("id")
	if err := DB.First(&pesanan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, pesanan)
}
