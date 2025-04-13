package models

import "gorm.io/gorm"

type Inventaris struct {
	gorm.Model
	ID       uint   `json:"id"`
	IDProduk uint   `json:"id_produk"`
	Jumlah   int    `json:"jumlah"`
	Lokasi   string `json:"lokasi"`
	Produk   Produk `gorm:"foreignKey:IDProduk"`
}
