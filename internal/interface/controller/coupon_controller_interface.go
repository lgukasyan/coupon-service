package controller

import "github.com/gin-gonic/gin"

type ICouponController interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Apply(*gin.Context)
}
