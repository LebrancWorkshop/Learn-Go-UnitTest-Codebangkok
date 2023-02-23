package handlers_test

import (
	"fmt"
	"io"
	"lebrancconvas/gounittest/handlers"
	"lebrancconvas/gounittest/services"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success handle", func(t *testing.T) {
		// Arrange 
		amount := 100
		expected := 80

		promoService := services.NewPromotionServicesMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotiionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
			// assert.Equal(t, "90", string(body))
		}
	})
}