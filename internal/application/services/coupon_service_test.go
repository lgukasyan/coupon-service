package services

import (
	"coupon_service/internal/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CouponRepositoryMock struct {
	mock.Mock
}

func (m *CouponRepositoryMock) FindByCode(code string) (bool, error) {
	args := m.Called(code)
	return args.Bool(0), args.Error(1)
}

func (m *CouponRepositoryMock) Create(coupon *model.Coupon) error {
	args := m.Called(coupon)
	return args.Error(0)
}

func TestCouponService(t *testing.T) {
	coupon := model.Coupon{
		Code:           "12345",
		Discount:       1,
		MinBasketValue: 100,
	}

	mockRepository := new(CouponRepositoryMock)
	service := NewCouponService(mockRepository)

	// Case 1
	t.Run("Coupon is nil", func(t *testing.T) {
		err := service.Create(nil)
		assert.EqualError(t, err, "coupon is empty")
	})

	// Case 2
	t.Run("Coupon code already exists", func(t *testing.T) {
		mockRepository.On("FindByCode", coupon.Code).Return(true, nil).Once()
		err := service.Create(&coupon)
		assert.EqualError(t, err, "coupon code already exists")
		mockRepository.AssertExpectations(t)
	})

	// Case 3
	t.Run("Cupon created successfully", func(t *testing.T) {
		mockRepository.On("FindByCode", coupon.Code).Return(false, nil).Once()
		mockRepository.On("Create", &coupon).Return(nil).Once()
		err := service.Create(&coupon)
		assert.NoError(t, err)
		mockRepository.AssertExpectations(t)
	})
}
