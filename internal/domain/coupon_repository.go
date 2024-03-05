package domain

import "coupon_service/internal/domain/model"

type ICouponRepository interface {
	FindByCode(string) (bool, error)
	Create(*model.Coupon) error
}
