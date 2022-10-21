package repositories

import (
	"nutech/models"

	"gorm.io/gorm"
)

type BarangRepository interface {
	FindBarangs() ([]models.Barang, error)
	GetBarang(ID int) (models.Barang, error)
	CreateBarang(barang models.Barang) (models.Barang, error)
	UpdateBarang(barang models.Barang) (models.Barang, error)
	DeleteBarang(barang models.Barang) (models.Barang, error)
}

func RepositoryBarang(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBarangs() ([]models.Barang, error) {
	var barangs []models.Barang
	err := r.db.Find(&barangs).Error

	return barangs, err
}

func (r *repository) GetBarang(ID int) (models.Barang, error) {
	var barang models.Barang
	err := r.db.First(&barang, ID).Error

	return barang, err
}

func (r *repository) CreateBarang(barang models.Barang) (models.Barang, error) {
	err := r.db.Preload("User").Create(&barang).Error

	return barang, err
}

func (r *repository) UpdateBarang(barang models.Barang) (models.Barang, error) {
	err := r.db.Save(&barang).Error

	return barang, err
}

func (r *repository) DeleteBarang(barang models.Barang) (models.Barang, error) {
	err := r.db.Delete(&barang).Error

	return barang, err
}
