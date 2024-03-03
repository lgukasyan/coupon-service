package infrastructure

import (
	"coupon_service/config"
	"coupon_service/internal/domain/model"
	"coupon_service/internal/infrastructure/db"
	"coupon_service/internal/infrastructure/router"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() (*gin.Engine, error) {
	// Create new env obj
	env := config.NewEnvConfig()

	// Load env
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	// Check DB connection and then start server
	db, err := db.New().Connect()
	if err != nil {
		return nil, err
	}

	// Migration
	if err := db.AutoMigrate(&model.Coupon{}); err != nil {
		return nil, err
	}

	// Create gin.engine
	r := gin.Default()
	if r == nil {
		return nil, errors.New("gin.engine is empty")
	}

	// Set up api group route
	apiGroup := r.Group("/api")

	// Set up coupon route
	router.SetUpCouponRouter(apiGroup, db)

	// Return gin.engine
	return r, nil
}
