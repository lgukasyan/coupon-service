package controller

import (
	"bytes"
	"coupon_service/internal/domain/model"
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

func TestCouponControllerCreateValidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(CouponServiceMock)
	controller := NewCouponController(mockService)

	r.POST("/coupon", controller.Create)
	mockService.On("Create", mock.AnythingOfType("*model.Coupon")).Return(nil)

	body := bytes.NewBufferString(`{"code": "12345", "discount": 2, "min_basket_value": 2}`)
	req, _ := http.NewRequest("POST", "/coupon", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestCouponControllerCreateInvalidKeyName(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(CouponServiceMock)

	controller := NewCouponController(mockService)
	r.POST("/coupon", controller.Create)

	body := bytes.NewBufferString(`{"wrong": "12345", "discount": 1, "min_basket_value": 1}`)
	req, err := http.NewRequest("POST", "/coupon", body)
	assert.NoError(t, err, "error sending the GET request")
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}

func TestCouponControllerCreateInvalidString(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(CouponServiceMock)

	controller := NewCouponController(mockService)
	r.POST("/coupon", controller.Create)

	body := bytes.NewBufferString(``)
	req, err := http.NewRequest("POST", "/coupon", body)
	assert.NoError(t, err, "error sending the GET request")
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}
