package db

import (
	"test-ITMX/internal/models"

	"github.com/stretchr/testify/mock"
)

type mockCustomersRepository struct {
	mock.Mock
}

func (mockCustomersRepo *mockCustomersRepository) Create(req models.CustomersRepository) error {
	args := mockCustomersRepo.Called()
	return args.Error(0)
}

// DeleteByID implements CustomersRepository.
func (mockCustomersRepo *mockCustomersRepository) DeleteByID(id int) error {
	args := mockCustomersRepo.Called()
	return args.Error(0)
}

// GetByID implements CustomersRepository.
func (mockCustomersRepo *mockCustomersRepository) GetByID(id int) (models.CustomersRepository, error) {
	args := mockCustomersRepo.Called()
	return args.Get(0).(models.CustomersRepository), args.Error(1)
}

// UpdateByID implements CustomersRepository.
func (mockCustomersRepo *mockCustomersRepository) UpdateByID(id int, req models.CustomersRepository) error {
	args := mockCustomersRepo.Called()
	return args.Error(0)
}
func NewCustomerRepositoryMock() *mockCustomersRepository {
	return &mockCustomersRepository{}
}
