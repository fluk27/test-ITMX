package handlers

import (
	"net/http"
	"test-ITMX/errs"

	"github.com/labstack/echo/v4"
)

type CustomersHandler interface {
	CreateCustomersHandler(c echo.Context) error
	GetCustomersByIdHandler(c echo.Context) error
	UpdateCustomersHandler(c echo.Context) error
	DeleteCustomersHandler(c echo.Context) error
}

func HandlerError(err error) *echo.HTTPError {
	switch e := err.(type) {
	case errs.AppError:
		return echo.NewHTTPError(e.Code, e.Message)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}
