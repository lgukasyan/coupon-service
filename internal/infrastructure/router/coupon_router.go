package router

import (
	"coupon_service/internal/application/services"
	"coupon_service/internal/interface/controller"
	"coupon_service/internal/interface/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpCouponRouter(r *gin.RouterGroup, db *gorm.DB) {
	// Initialize
	repository := repository.NewCouponRepository(db)
	service := services.NewCouponService(repository)
	controller := controller.NewCouponController(service)

	// Routes
	cGroup := r.Group("/coupon")
	cGroup.GET("/", controller.Ping)
	cGroup.POST("/create", controller.Create)
}
