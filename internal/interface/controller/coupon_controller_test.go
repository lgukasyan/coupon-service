package controller

import (
	"bytes"
	"coupon_service/internal/domain/model"
	"coupon_service/internal/interface/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CouponServiceMock struct {
	mock.Mock
}

func (m *CouponServiceMock) Create(coupon *model.Coupon) error {
	args := m.Called(coupon)
	return args.Error(0)
}

func (m *CouponServiceMock) Get() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *CouponServiceMock) Apply(basket *dto.BasketRequestDTO) (*dto.BasketResponseDTO, error) {
	args := m.Called(basket)
	return args.Get(0).(*dto.BasketResponseDTO), args.Error(1)
}

func TestCouponController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(CouponServiceMock)
	controller := NewCouponController(mockService)

	r := gin.Default()
	r.GET("/coupon", controller.Get)
	r.POST("/coupon/create", controller.Create)
	r.POST("/coupon/apply", controller.Apply)

	t.Run("Valid coupon schema, should return 200", func(t *testing.T) {
		mockService.On("Create", mock.AnythingOfType("*model.Coupon")).Return(nil)
		body := bytes.NewBufferString(`{"code": "DESC1", "discount": 2, "min_basket_value": 2}`)

		req, _ := http.NewRequest("POST", "/coupon/create", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Invalid coupon schema, should return 400", func(t *testing.T) {
		mockService.On("Create", mock.AnythingOfType("*model.Coupon")).Return(nil)
		body := bytes.NewBufferString(`{"": "DESC1", "discount": 2, "min_basket_value": 2}`)

		req, _ := http.NewRequest("POST", "/coupon/create", body)
		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Get all codes, should return 200", func(t *testing.T) {
		mockService.On("Get").Return([]string{}, nil)
		req, _ := http.NewRequest("GET", "/coupon", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Apply coupon valid schema, should return 200", func(t *testing.T) {
		mockService.On("Apply", mock.Anything).Return(&dto.BasketResponseDTO{}, nil)
		body := bytes.NewBufferString(`{"code": "DESC1", "value": 100}`)
		req, _ := http.NewRequest("POST", "/coupon/apply", body)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Apply coupon invalid schema, should return 200", func(t *testing.T) {
		mockService.On("Apply", mock.Anything).Return(&dto.BasketResponseDTO{}, nil)
		body := bytes.NewBufferString(`{"coe": "DESC1", "value": 100}`)
		req, _ := http.NewRequest("POST", "/coupon/apply", body)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})
}
