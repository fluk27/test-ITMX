package services_test

import (
	"errors"
	"test-ITMX/config"
	"test-ITMX/constant"
	"test-ITMX/internal/models"
	"test-ITMX/internal/repositories/db"
	"test-ITMX/internal/services"
	"test-ITMX/loggers"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetCustomer(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		mockData      models.CustomersRepository
		expectSuccess models.CustomerResponse
		expectError   error
	}{
		{
			name:      "TestGetCustomerSuccess",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   1,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerGetSuccessMessage,
				Data: &models.CustomersData{
					ID:   1,
					Name: "testCustomer",
					Age:  16,
				},
			},
			expectError: nil,
		},
		{
			name:      "TestGetCustomerErrorInternalServerError",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   1,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorMessageInternalServerError),
		},
		{
			name:      "TestGetCustomerFindNotFound",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   0,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorsMessageFindNotFound),
		},
		{
			name:      "TestGetCustomerNotMatchFindNotFound",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   0,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			customerRepo := db.NewCustomerRepositoryMock()

			switch tC.name {
			case "TestGetCustomerFindNotFound":
				customerRepo.On("GetByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				break
			case "TestGetCustomerNotMatchFindNotFound":
				customerRepo.On("GetByID").Return(tC.mockData, errors.New(""))

				break
			case "TestGetCustomerErrorInternalServerError":
				customerRepo.On("GetByID").Return(tC.mockData, errors.New(""))

				break
			default:
				customerRepo.On("GetByID").Return(tC.mockData, nil)
				break
			}

			customerSvc := services.NewCustomersService(customerRepo)

			resp, err := customerSvc.GetByID(tC.requestId)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestCreateCustomer(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		request       models.CustomerRequest
		mockData      models.CustomersRepository
		expectSuccess models.CustomerResponse
		expectError   error
	}{
		{
			name: "createCustomerSuccess",
			request: models.CustomerRequest{
				Name: "testCustomer",
				Age:  16,
			},
			mockData: models.CustomersRepository{
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name: "createCustomerError",
			request: models.CustomerRequest{
				Name: "testCustomer",
				Age:  16,
			},
			mockData: models.CustomersRepository{
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			customerRepo := db.NewCustomerRepositoryMock()
			// customerRepo.On("GetByID").Return(tC.mockdata, nil)
			if tC.expectError != nil {
				customerRepo.On("Create").Return(errors.New(constant.CustomerErrorMessageInternalServerError))
			} else {
				customerRepo.On("Create").Return(nil)
			}

			customerSvc := services.NewCustomersService(customerRepo)
			resp, err := customerSvc.Create(tC.request)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestUpdateCustomer(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		requestBody   models.CustomerRequest
		mockData      models.CustomersRepository
		expectSuccess models.CustomerResponse
		expectError   error
	}{
		{
			name:      "TestUpdateCustomerSuccess",
			requestId: 1,
			requestBody: models.CustomerRequest{
				Name: "testCustomer",
				Age:  16,
			},
			mockData: models.CustomersRepository{
				ID:   1,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerUpdateSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name:      "TestUpdateCustomerErrorInternalServerError",
			requestId: 1,
			requestBody: models.CustomerRequest{
				Name: "testCustomer",
				Age:  16,
			},
			mockData: models.CustomersRepository{
				ID:   1,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorMessageInternalServerError),
		},
		{
			name:      "TestUpdateCustomerFindNotFound",
			requestId: 1,
			requestBody: models.CustomerRequest{
				Name: "testCustomer",
				Age:  16,
			},
			mockData: models.CustomersRepository{
				ID:   0,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorsMessageFindNotFound),
		},
		{
			name:      "TestUpdateCustomerNotMatchFindNotFound",
			requestId: 1,
			requestBody: models.CustomerRequest{
				Name: "testCustomer",
				Age:  16,
			},
			mockData: models.CustomersRepository{
				ID:   0,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			customerRepo := db.NewCustomerRepositoryMock()

			switch tC.name {
			case "TestUpdateCustomerFindNotFound":
				customerRepo.On("GetByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				customerRepo.On("UpdateByID").Return(tC.expectError)
				break
			case "TestUpdateCustomerNotMatchFindNotFound":
				customerRepo.On("GetByID").Return(tC.mockData, errors.New(""))
				customerRepo.On("UpdateByID").Return(tC.expectError)
				break
			case "TestUpdateCustomerErrorInternalServerError":
				customerRepo.On("GetByID").Return(tC.mockData, nil)
				customerRepo.On("UpdateByID").Return(tC.expectError)
				break
			default:
				customerRepo.On("GetByID").Return(tC.mockData, nil)
				customerRepo.On("UpdateByID").Return(nil)
				break
			}

			customerSvc := services.NewCustomersService(customerRepo)

			resp, err := customerSvc.UpdateByID(tC.requestId, tC.requestBody)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}

func TestDeleteCustomer(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		mockData      models.CustomersRepository
		expectSuccess models.CustomerResponse
		expectError   error
	}{
		{
			name:      "TestDeleteCustomerSuccess",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   1,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerDeleteSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name:      "TestDeleteCustomerErrorInternalServerError",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   1,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorMessageInternalServerError),
		},
		{
			name:      "TestDeleteCustomerFindNotFound",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   0,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorsMessageFindNotFound),
		},
		{
			name:      "TestDeleteCustomerNotMatchFindNotFound",
			requestId: 1,
			mockData: models.CustomersRepository{
				ID:   0,
				Name: "testCustomer",
				Age:  16,
			},

			expectSuccess: models.CustomerResponse{
				Message: constant.CustomerCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.CustomerErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			customerRepo := db.NewCustomerRepositoryMock()

			switch tC.name {
			case "TestDeleteCustomerFindNotFound":
				customerRepo.On("GetByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				customerRepo.On("DeleteByID").Return(tC.expectError)
				break
			case "TestDeleteCustomerNotMatchFindNotFound":
				customerRepo.On("GetByID").Return(tC.mockData, errors.New(""))
				customerRepo.On("DeleteByID").Return(tC.expectError)
				break
			case "TestDeleteCustomerErrorInternalServerError":
				customerRepo.On("GetByID").Return(tC.mockData, nil)
				customerRepo.On("DeleteByID").Return(tC.expectError)
				break
			default:
				customerRepo.On("GetByID").Return(tC.mockData, nil)
				customerRepo.On("DeleteByID").Return(nil)
				break
			}

			customerSvc := services.NewCustomersService(customerRepo)

			resp, err := customerSvc.DeleteByID(tC.requestId)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
