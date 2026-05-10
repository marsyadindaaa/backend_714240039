package main

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	IDDonasi int `gorm:"primaryKey;column:id_donasi" json:"id_donasi"`
	NamaDonatur string `json:"nama_donatur"`
	Nominal int `json:"nominal"`
	TanggalDonasi time.Time `json:"tanggal_donasi"`
	MetodePembayaran string `json:"metode_pembayaran"`
	Status string `json:"status"`
}

func MigrateDonations(db *gorm.DB){
	db.AutoMigrate(&Donation{})
}