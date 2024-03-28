package models

type CustomersRepository struct {
	ID   int
	Name string
	Age  int
}

func (CustomersRepository) TableName() string {
	return "customers"
}
