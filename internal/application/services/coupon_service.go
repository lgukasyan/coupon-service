package services

import "coupon_service/internal/domain"

type CouponService struct {
	couponRepository domain.ICouponRepository
}

func NewCouponService(repo domain.ICouponRepository) ICouponService {
	return &CouponService{
		couponRepository: repo,
	}
}
