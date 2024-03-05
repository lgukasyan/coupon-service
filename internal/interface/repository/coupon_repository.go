package repository

import (
	"coupon_service/internal/domain"
	"coupon_service/internal/domain/model"
	"errors"

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

// Find By Code (PrimaryKey)
func (repo *CouponRepository) FindByCode(code string) (bool, error) {
	var coupon model.Coupon
	err := repo.DB.First(&coupon, code).Error

	if errors.Is(gorm.ErrRecordNotFound, err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Create
func (repo *CouponRepository) Create(coupon *model.Coupon) error {
	return repo.DB.Create(&coupon).Error
}
