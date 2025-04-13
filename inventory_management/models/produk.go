package models

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	ID        uint    `json:"id"`
	Nama      string  `json:"nama"`
	Deskripsi string  `json:"deskripsi"`
	Harga     float64 `json:"harga"`
	Kategori  string  `json:"kategori"`
}
