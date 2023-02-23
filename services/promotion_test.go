package services_test

import (
	"lebrancconvas/gounittest/repositories"
	"lebrancconvas/gounittest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	type testCase struct {
		name string
		purchaseMin int
		discountPercent int
		amount int
		expected int
	}

	cases := []testCase{
		{name: "PurchaseMin 100, DiscountPercent 20, Amount 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "PurchaseMin 100, DiscountPercent 20, Amount 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
		{name: "PurchaseMin 100, DiscountPercent 20, Amount 0", purchaseMin: 100, discountPercent: 20, amount: 0, expected: 0},
		{name: "PurchaseMin 100, DiscountPercent 20, Amount 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arrange (เตรียมของ)
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID: 1,
				PurchaseMin: c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)
		
			promoService := services.NewPromotionService(promoRepo)
		
			// Act (ทำจริง)
			discount, _ := promoService.CalculateDiscount(c.amount)
			expected := c.expected
		
			// Assert (ตรวจสอบ)
			assert.Equal(t, expected, discount)
		})
	}



}