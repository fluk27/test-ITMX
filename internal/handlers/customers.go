package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"test-ITMX/internal/models"
	"test-ITMX/internal/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type customersHandler struct {
	customersService services.CustomersService
}

// CreateCustomersHandler implements CustomersHandler.
func (customersHandle customersHandler) CreateCustomersHandler(c echo.Context) error {
	customerReq := new(models.CustomerRequest)
	if err := c.Bind(customerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(customerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	customerResp, err := customersHandle.customersService.Create(*customerReq)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusCreated, customerResp, "")
}

// DeleteCustomersHandler implements CustomersHandler.
func (customersHandle customersHandler) DeleteCustomersHandler(c echo.Context) error {
	paramsId := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(paramsId)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	customerResp, err := customersHandle.customersService.DeleteByID(id)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, customerResp, "")
}

// GetCustomersByIdHandler implements CustomersHandler.
func (customersHandle customersHandler) GetCustomersByIdHandler(c echo.Context) error {

	Id := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(Id)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}
	id, err := strconv.Atoi(Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	customerResp, err := customersHandle.customersService.GetByID(id)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, customerResp, "")
}

// UpdateCustomersHandler implements CustomersHandler.
func (customersHandle customersHandler) UpdateCustomersHandler(c echo.Context) error {
	paramsId := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(paramsId)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	customerReq := new(models.CustomerRequest)
	if err = c.Bind(customerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(customerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	customerResp, err := customersHandle.customersService.UpdateByID(id, *customerReq)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, customerResp, "")
}

func NewCustomersHandler(customersService services.CustomersService) CustomersHandler {
	return customersHandler{customersService: customersService}
}
