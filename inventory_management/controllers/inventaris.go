package controllers

import (
	"inventory_management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LihatStok(c *gin.Context) {
	var inventaris models.Inventaris
	idProduk := c.Param("id_produk")
	if err := DB.First(&inventaris, "id_produk = ?", idProduk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan di inventaris"})
		return
	}
	c.JSON(http.StatusOK, inventaris)
}

func UpdateStok(c *gin.Context) {
	var inventaris models.Inventaris
	idProduk := c.Param("id_produk")

	if err := DB.First(&inventaris, "id_produk = ?", idProduk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan di inventaris"})
		return
	}
	if err := c.ShouldBindJSON(&inventaris); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&inventaris)
	c.JSON(http.StatusOK, inventaris)
}
