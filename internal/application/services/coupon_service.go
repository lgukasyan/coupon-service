package services

import (
	"coupon_service/internal/domain"
	"coupon_service/internal/domain/model"
	"coupon_service/internal/interface/dto"
	"errors"
)

type CouponService struct {
	couponRepository domain.ICouponRepository
}

func NewCouponService(repo domain.ICouponRepository) ICouponService {
	return &CouponService{
		couponRepository: repo,
	}
}

func (s *CouponService) Create(coupon *model.Coupon) error {
	// Check if pointer is nil
	if coupon == nil {
		return errors.New("coupon is empty")
	}

	// Check if coupon code exists
	ok, err := s.couponRepository.Exists(coupon.Code)
	if err != nil {
		return err
	}

	// If exists, return error
	if ok {
		return errors.New("coupon code already exists")
	}

	// Create
	if err := s.couponRepository.Create(coupon); err != nil {
		return errors.New("error while creating a coupon")
	}

	return nil
}

func (s *CouponService) Get() ([]string, error) {
	codes, err := s.couponRepository.Get()
	if err != nil {
		return codes, err
	}

	return codes, nil
}

func (s *CouponService) Apply(basket *dto.BasketRequestDTO) (*dto.BasketResponseDTO, error) {
	coupon, err := s.couponRepository.FindByCode(basket.Code)
	if err != nil {
		return nil, err
	}

	if coupon == nil {
		return nil, errors.New("coupon is nil")
	}

	resp := &dto.BasketResponseDTO{
		Value:           basket.Value,
		AppliedDiscount: coupon.Discount,
	}

	// Rules
	if basket.Value < coupon.MinBasketValue {
		resp.ApplicationSuccessful = false
		return resp, nil
	}

	finalValue := basket.Value - ((basket.Value * coupon.Discount) / 100)
	resp.ApplicationSuccessful = true
	resp.FinalValue = &finalValue

	return resp, nil
}
