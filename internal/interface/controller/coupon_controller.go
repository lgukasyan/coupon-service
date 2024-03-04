package controller

import (
	"coupon_service/internal/application/services"
	"coupon_service/internal/domain/model"
	"coupon_service/internal/interface/response"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CouponController struct {
	couponService services.ICouponService
}

func NewCouponController(cs services.ICouponService) ICouponController {
	return &CouponController{}
}

func (c *CouponController) Ping(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
}

func (c *CouponController) Create(ctx *gin.Context) {
	var coupon model.Coupon
	if err := json.NewDecoder(ctx.Request.Body).Decode(&coupon); err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(coupon); err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.JSON(ctx, http.StatusCreated, nil)
}
