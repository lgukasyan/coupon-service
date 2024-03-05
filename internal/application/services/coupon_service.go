package services

import (
	"coupon_service/internal/domain"
	"coupon_service/internal/domain/model"
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
	ok, err := s.couponRepository.FindByCode(coupon.Code)
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
