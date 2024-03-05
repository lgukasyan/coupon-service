package services

import (
	"coupon_service/internal/domain/model"
	"coupon_service/internal/interface/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CouponRepositoryMock struct {
	mock.Mock
}

func (m *CouponRepositoryMock) FindByCode(code string) (*model.Coupon, error) {
	args := m.Called(code)

	result, ok := args.Get(0).(*model.Coupon)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

func (m *CouponRepositoryMock) Exists(code string) (bool, error) {
	args := m.Called(code)
	return args.Bool(0), args.Error(1)
}

func (m *CouponRepositoryMock) Create(coupon *model.Coupon) error {
	args := m.Called(coupon)
	return args.Error(0)
}

func (m *CouponRepositoryMock) Get() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func TestCouponService(t *testing.T) {
	coupon := model.Coupon{
		Code:           "DESC1",
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
		mockRepository.On("Exists", coupon.Code).Return(true, nil).Once()
		err := service.Create(&coupon)
		assert.EqualError(t, err, "coupon code already exists")
		mockRepository.AssertExpectations(t)
	})

	// Case 3
	t.Run("Cupon created successfully", func(t *testing.T) {
		mockRepository.On("Exists", coupon.Code).Return(false, nil).Once()
		mockRepository.On("Create", &coupon).Return(nil).Once()
		err := service.Create(&coupon)
		assert.NoError(t, err)
		mockRepository.AssertExpectations(t)
	})

	// Case 4
	t.Run("Get coupons should be empty", func(t *testing.T) {
		mockRepository.On("Get").Return([]string{}, nil).Once()
		codes, err := service.Get()
		assert.NoError(t, err)
		assert.Empty(t, codes)
		mockRepository.AssertExpectations(t)
	})

	// Case 4
	t.Run("Valid coupon Apply", func(t *testing.T) {
		mockRepository.On("FindByCode", mock.Anything).Return(&model.Coupon{
			Code:           "DESC1",
			Discount:       1,
			MinBasketValue: 100,
		}, nil).Once()

		basketResponse, err := service.Apply(&dto.BasketRequestDTO{
			Code:  "DESC1",
			Value: 1000,
		})

		var finalValue uint = 990

		assert.Equal(t, &dto.BasketResponseDTO{
			Value:                 1000,
			AppliedDiscount:       1,
			ApplicationSuccessful: true,
			FinalValue:            &finalValue,
		}, basketResponse)

		assert.NoError(t, nil, err)
		mockRepository.AssertExpectations(t)
	})

	// Case 5
	t.Run("Invalid coupon Apply", func(t *testing.T) {
		mockRepository.On("FindByCode", mock.Anything).Return(&model.Coupon{
			Code:           "DESC1",
			Discount:       1,
			MinBasketValue: 100,
		}, nil).Once()

		basketResponse, err := service.Apply(&dto.BasketRequestDTO{
			Code:  "DESC1",
			Value: 99,
		})

		assert.Equal(t, &dto.BasketResponseDTO{
			Value:                 99,
			AppliedDiscount:       1,
			ApplicationSuccessful: false,
		}, basketResponse)

		assert.NoError(t, nil, err)
		mockRepository.AssertExpectations(t)
	})

	// Case 6
	t.Run("Invalid basket value", func(t *testing.T) {
		mockRepository.On("FindByCode", mock.Anything).Return(nil, nil).Once()
		basketResponse, err := service.Apply(&dto.BasketRequestDTO{
			Code:  "DESC1",
			Value: 99,
		})

		assert.Nil(t, basketResponse)
		assert.NoError(t, nil, err)
		mockRepository.AssertExpectations(t)
	})
}
