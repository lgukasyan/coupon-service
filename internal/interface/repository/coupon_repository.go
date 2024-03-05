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

// Exists?
func (repo *CouponRepository) Exists(code string) (bool, error) {
	var coupon model.Coupon
	err := repo.DB.First(&coupon, "code = ?", code).Error

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
	if coupon == nil {
		return errors.New("coupon is empty")
	}
	return repo.DB.Create(&coupon).Error
}

// Get all codes
func (repo *CouponRepository) Get() ([]string, error) {
	var codes []string
	if err := repo.DB.Model(&model.Coupon{}).Select("code").Find(&codes).Error; err != nil {
		return codes, err
	}

	return codes, nil
}

// Get coupon
func (repo *CouponRepository) FindByCode(code string) (*model.Coupon, error) {
	var coupon model.Coupon
	err := repo.DB.First(&coupon, "code = ?", code).Error

	if err != nil {
		return nil, err
	}

	return &coupon, nil
}
