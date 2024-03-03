package infrastructure

import (
	"coupon_service/internal/domain/model"
	"coupon_service/internal/infrastructure/db"
	"coupon_service/internal/infrastructure/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	// Check DB connection and then start server
	db, err := db.New().Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Migration
	if err := db.AutoMigrate(&model.Coupon{}); err != nil {
		log.Fatal(err)
	}

	// Create gin.engine
	r := gin.Default()
	if r == nil {
		log.Fatal("gin.engine empty")
	}

	// Set up api group route
	apiGroup := r.Group("/api")

	// Set up coupon route
	router.SetUpCouponRouter(apiGroup, db)

	// Read port env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port env is empty")
	}

	// Start
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
