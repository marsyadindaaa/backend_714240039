package main

import "gorm.io/gorm"

type DonasiRepository struct {
	DB *gorm.DB
}

func NewDonasiRepository(db *gorm.DB) *DonasiRepository{
	return &DonasiRepository{DB: db}
}

func (r *DonasiRepository) GetAllDonations() ([]Donation, error){
	var donations []Donation
	err := r.DB.Find(&donations).Error
	return donations, err
}

func (r *DonasiRepository) GetDonationByID(id_donasi int) (Donation, error){
	var donation Donation
	err := r.DB.First(&donation,"id_donasi = ?", id_donasi).Error
	return donation, err
}

func (r *DonasiRepository) CreateDonation(donation *Donation) error{
	return r.DB.Create(donation).Error
}

func (r *DonasiRepository) UpdateDonation(donation *Donation) error{
	return r.DB.Model(donation).Updates(map[string]interface{}{
		"nama_donatur": donation.NamaDonatur,
		"nominal": donation.Nominal,
		"tanggal_donasi": donation.TanggalDonasi,
		"metode_pembayaran": donation.MetodePembayaran,
		"status": donation.Status,
	}).Error
}

func (r *DonasiRepository) DeleteDonation(id_donasi int) error{
	return r.DB.Delete(&Donation{}, "id_donasi = ?", id_donasi).Error
}


