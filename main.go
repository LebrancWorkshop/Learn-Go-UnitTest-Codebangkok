package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type CustomerRepository interface {
	GetCustomer(id int) (name string, age int, err error)
	Hello()
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0)	, args.Int(1), args.Error(2)
}

func (m *CustomerRepositoryMock) Hello() {
	fmt.Println(m)
}

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Nagi", 90, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	name, age, err := c.GetCustomer(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("--Profile--\nName: %v\nAge: %v\n", name, age)
}