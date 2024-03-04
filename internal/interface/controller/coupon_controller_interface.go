package controller

import "github.com/gin-gonic/gin"

type ICouponController interface {
	Ping(*gin.Context)
	Create(*gin.Context)
}
