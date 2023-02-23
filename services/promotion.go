package services

import "lebrancconvas/gounittest/repositories"

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	promoRepo repositories.PromotionRepository
}

func NewPromotionService(promoRepo repositories.PromotionRepository) PromotionService {
	return promotionService{promoRepo: promoRepo}
}

func (s promotionService) CalculateDiscount(amount int) (int, error) {
	if amount <= 0 {
		return 0, ErrNotValid
	}

	promotion, err := s.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promotion.PurchaseMin {
		return amount - (amount * promotion.DiscountPercent / 100), nil
	}

	return amount, nil
}

