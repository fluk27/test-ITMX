package db

import (
	"test-ITMX/internal/models"

	"gorm.io/gorm"
)

type customersRepository struct {
	db *gorm.DB
}

// Create implements CustomersRepository.
func (c customersRepository) Create(req models.CustomersRepository) error {
	if err := c.db.Create(&req).Error; err != nil {
		return err
	}

	return nil
}

// DeleteByID implements CustomersRepository.
func (c customersRepository) DeleteByID(id int) error {
	if err := c.db.Where("id", id).Delete(&models.CustomersRepository{}).Error; err != nil {
		return err
	}

	return nil
}

// GetByID implements CustomersRepository.
func (c customersRepository) GetByID(id int) (models.CustomersRepository, error) {
	customerRepoResp := models.CustomersRepository{}
	if err := c.db.Where("id", id).First(&customerRepoResp).Error; err != nil {
		return customerRepoResp, err
	}

	return customerRepoResp, nil
}

// UpdateByID implements CustomersRepository.
func (c customersRepository) UpdateByID(id int, req models.CustomersRepository) error {
	if err := c.db.Where("id", id).Updates(&req).Error; err != nil {
		return err
	}

	return nil
}

func NewCustomersRepository(db *gorm.DB) CustomersRepository {
	return customersRepository{db: db}
}
