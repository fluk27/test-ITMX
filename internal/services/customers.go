package services

import (
	"errors"
	"fmt"
	"test-ITMX/constant"
	"test-ITMX/errs"
	"test-ITMX/internal/models"
	"test-ITMX/internal/repositories/db"
	"test-ITMX/loggers"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type customersService struct {
	customersRepo db.CustomersRepository
}

// Create implements CustomersService.
func (c customersService) Create(req models.CustomerRequest) (models.CustomerResponse, error) {
	customerRepReq := models.CustomersRepository{
		Name: req.Name,
		Age:  req.Age,
	}
	err := c.customersRepo.Create(customerRepReq)
	if err != nil {
		loggers.Error(fmt.Sprintf("Create=%v", err.Error()),
			zap.String("type", "repo"),
			zap.Any("customerRepReq", customerRepReq),
			zap.Error(err))
		return models.CustomerResponse{}, errs.NewInternalServerError(constant.CustomerErrorMessageInternalServerError)
	}
	return models.CustomerResponse{
		Message: constant.CustomerCreateSuccessMessage,
	}, nil
}

// DeleteByID implements CustomersService.
func (c customersService) DeleteByID(id int) (models.CustomerResponse, error) {
	_, err := c.customersRepo.GetByID(id)
	if err != nil {
		loggers.Error(fmt.Sprintf("GetByID=%v", err.Error()),
			zap.String("type", "repo"),
			zap.Int("id", id),
			zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.CustomerResponse{}, errs.NewNotFoundError(constant.CustomerErrorsMessageFindNotFound)
		} else {

			return models.CustomerResponse{}, errs.NewInternalServerError(constant.CustomerErrorMessageInternalServerError)
		}
	}

	err = c.customersRepo.DeleteByID(id)
	if err != nil {
		loggers.Error(fmt.Sprintf("DeleteByID=%v", err.Error()),
			zap.String("type", "repo"),
			zap.Int("id", id),
			zap.Error(err))
		return models.CustomerResponse{}, errs.NewInternalServerError(constant.CustomerErrorMessageInternalServerError)
	}
	return models.CustomerResponse{

		Message: constant.CustomerDeleteSuccessMessage,
	}, nil
}

// GetByID implements CustomersService.
func (c customersService) GetByID(id int) (models.CustomerResponse, error) {
	customer, err := c.customersRepo.GetByID(id)
	if err != nil {
		loggers.Error(fmt.Sprintf("GetByID=%v", err.Error()),
			zap.String("type", "repo"),
			zap.Int("id", id),
			zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.CustomerResponse{}, errs.NewInternalServerError(constant.CustomerErrorsMessageFindNotFound)
		} else {

			return models.CustomerResponse{}, errs.NewInternalServerError(constant.CustomerErrorMessageInternalServerError)
		}
	}
	customerdata := models.CustomersData{
		ID:   customer.ID,
		Name: customer.Name,
		Age:  customer.Age,
	}
	return models.CustomerResponse{

		Message: constant.CustomerGetSuccessMessage,
		Data:    &customerdata,
	}, nil
}

// UpdateByID implements CustomersService.
func (c customersService) UpdateByID(id int, req models.CustomerRequest) (models.CustomerResponse, error) {
	_, err := c.customersRepo.GetByID(id)
	if err != nil {
		loggers.Error(fmt.Sprintf("GetByID=%v", err.Error()),
			zap.String("type", "repo"),
			zap.Int("id", id),
			zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.CustomerResponse{}, errs.NewNotFoundError(constant.CustomerErrorsMessageFindNotFound)
		} else {

			return models.CustomerResponse{}, errs.NewInternalServerError(constant.CustomerErrorMessageInternalServerError)
		}
	}
	customerRepReq := models.CustomersRepository{
		Name: req.Name,
		Age:  req.Age,
	}
	err = c.customersRepo.UpdateByID(id, customerRepReq)
	if err != nil {
		loggers.Error(fmt.Sprintf("UpdateByID=%v", err.Error()),
			zap.String("type", "repo"),
			zap.Int("id", id),
			zap.Error(err))
		return models.CustomerResponse{}, errs.NewInternalServerError(constant.CustomerErrorMessageInternalServerError)
	}
	return models.CustomerResponse{

		Message: constant.CustomerUpdateSuccessMessage,
	}, nil
}

func NewCustomersService(customersRepo db.CustomersRepository) CustomersService {
	return customersService{customersRepo: customersRepo}
}
