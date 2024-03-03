package infrastructure

import (
	"coupon_service/internal/infrastructure/db"
	"coupon_service/internal/infrastructure/router"
	"log"
	"os"
)

func Start() {
	// Check DB connection and then start server
	_, err := db.New().Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Set up API router
	r := router.SetUpRouter()
	if r == nil {
		log.Fatal("gin.engine is empty")
	}

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
