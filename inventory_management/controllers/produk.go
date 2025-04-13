package controllers

import (
	"inventory_management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TambahProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&produk)
	c.JSON(http.StatusOK, produk)
}

func LihatSemuaProduk(c *gin.Context) {
	var produk []models.Produk
	if err := DB.Find(&produk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func LihatProdukByID(c *gin.Context) {
	var produk models.Produk
	id := c.Param("id")
	if err := DB.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func LihatProdukByKategori(c *gin.Context) {
	var produk []models.Produk
	kategori := c.Param("kategori")
	if err := DB.Where("kategori = ?", kategori).Find(&produk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func UpdateProduk(c *gin.Context) {
	var produk models.Produk
	id := c.Param("id")
	if err := DB.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&produk)
	c.JSON(http.StatusOK, produk)
}

func HapusProduk(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&models.Produk{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}
