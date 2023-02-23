package handlers

import (
	"lebrancconvas/gounittest/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoService services.PromotionService
}

func NewPromotiionHandler(promoService services.PromotionService) PromotionHandler {
	return promotionHandler{promoService: promoService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	amount := c.Query("amount")
	amountNumber, err := strconv.Atoi(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	discount, err := h.promoService.CalculateDiscount(amountNumber)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendString(strconv.Itoa(discount))
}
