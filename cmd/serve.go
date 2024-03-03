package cmd

import (
	"coupon_service/config"
	"coupon_service/internal/infrastructure"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start coupon service",
	Run: func(cmd *cobra.Command, args []string) {
		// Create new env obj
		env := config.New()

		// Load env
		if err := env.Load(); err != nil {
			log.Fatal(err)
		}

		// Get env port
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("env port is empty")
		}

		// Pass port and start server
		infrastructure.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
