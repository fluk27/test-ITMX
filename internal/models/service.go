package models

type CustomerResponse struct {
	Message string         `json:"message"`
	Data    *CustomersData `json:"data,omitempty"`
}
type CustomerRequest struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}
type CustomersData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
