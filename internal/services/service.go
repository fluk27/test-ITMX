package services

import (
	"test-ITMX/internal/models"
)

type CustomersService interface {
	GetByID(id int) (models.CustomerResponse, error)
	UpdateByID(id int, req models.CustomerRequest) (models.CustomerResponse, error)
	DeleteByID(id int) (models.CustomerResponse, error)
	Create(req models.CustomerRequest) (models.CustomerResponse, error)
}
