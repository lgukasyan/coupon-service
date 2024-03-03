package router

import (
	"coupon_service/internal/application/services"
	"coupon_service/internal/interface/controller"
	"coupon_service/internal/interface/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpCouponRouter(r *gin.RouterGroup, db *gorm.DB) {
	repository := repository.NewCouponRepository(db)
	service := services.NewCouponService(repository)
	controller := controller.NewCouponController(service)

	cGroup := r.Group("/coupon")
	cGroup.GET("/", controller.Ping)
}
