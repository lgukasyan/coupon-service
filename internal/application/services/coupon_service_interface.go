package services

import (
	"coupon_service/internal/domain/model"
	"coupon_service/internal/interface/dto"
)

type ICouponService interface {
	Create(*model.Coupon) error
	Get() ([]string, error)
	Apply(*dto.BasketRequestDTO) (*dto.BasketResponseDTO, error)
}
