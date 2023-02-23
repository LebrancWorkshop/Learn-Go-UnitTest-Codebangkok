package services

import "github.com/stretchr/testify/mock"

type promotionServicesMock struct {
	mock.Mock
}

func NewPromotionServicesMock() *promotionServicesMock {
	 return &promotionServicesMock{}
}

func (m *promotionServicesMock) CalculateDiscount(amount int) (int, error) {
	args := m.Called(amount)
	return args.Int(0), args.Error(1)
}

