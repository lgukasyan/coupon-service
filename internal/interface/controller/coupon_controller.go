package controller

import (
	"coupon_service/internal/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
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
