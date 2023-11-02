package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Bond", 18, nil)                //create mocking db
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found")) //create mocking db

	name, age, err := c.GetCustomer(1) //example for making query
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)

}

type CustomerRepositoryMock struct {
	mock.Mock
}

// query mnocking
func (m *CustomerRepositoryMock) GetCustomer(id int) (string, int, error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2) //return value according to index order
}
