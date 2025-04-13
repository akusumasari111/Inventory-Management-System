package models

import "gorm.io/gorm"

type Pesanan struct {
	gorm.Model
	ID       uint   `json:"id"`
	IDProduk uint   `json:"id_produk"`
	Jumlah   int    `json:"jumlah"`
	Tanggal  string `json:"tanggal_pesanan"`
	Produk   Produk `gorm:"foreignKey:IDProduk"`
}
