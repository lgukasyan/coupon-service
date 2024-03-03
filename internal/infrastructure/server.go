package infrastructure

import (
	"coupon_service/internal/infrastructure/router"
	"log"
)

func Start(port string) {
	r := router.SetUpRouter()
	if r == nil {
		log.Fatal("gin.engine is empty")
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
