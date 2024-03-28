package routers

import (
	"test-ITMX/internal/handlers"
	"test-ITMX/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(productPurchaseSvc services.CustomersService) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	//customers
	customersHandle := handlers.NewCustomersHandler(productPurchaseSvc)
	api := e.Group("/customers")
	api.POST("", customersHandle.CreateCustomersHandler)
	api.GET("/:id", customersHandle.GetCustomersByIdHandler)
	api.PUT("/:id", customersHandle.UpdateCustomersHandler)
	api.DELETE("/:id", customersHandle.DeleteCustomersHandler)
	return e
}
