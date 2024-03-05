package services

import "coupon_service/internal/domain/model"

type ICouponService interface {
	Create(*model.Coupon) error
	Get() ([]string, error)
}
