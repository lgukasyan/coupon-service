package cmd

import (
	"coupon_service/internal/infrastructure"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start coupon service",
	Run: func(cmd *cobra.Command, args []string) {
		// Get gin.engine
		r, err := infrastructure.Start()
		if err != nil {
			log.Fatal(err)
		}

		// Read port env
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("port env is empty")
		}

		// Start server
		if err := r.Run(":" + port); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
