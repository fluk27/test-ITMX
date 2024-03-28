package db

import (
	"test-ITMX/internal/models"
)

type CustomersRepository interface {
	GetByID(id int) (models.CustomersRepository, error)
	UpdateByID(id int, req models.CustomersRepository) error
	DeleteByID(id int) error
	Create(req models.CustomersRepository) error
}
