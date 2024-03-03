package repository

import (
	"coupon_service/internal/domain"

	"gorm.io/gorm"
)

type CouponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) domain.ICouponRepository {
	return &CouponRepository{
		DB: db,
	}
}
