package domain

import "coupon_service/internal/domain/model"

type ICouponRepository interface {
	Exists(string) (bool, error)
	FindByCode(string) (*model.Coupon, error)
	Create(*model.Coupon) error
	Get() ([]string, error)
}
