package cmd

import (
	"coupon_service/config"
	"coupon_service/internal/infrastructure"
	"log"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start coupon service",
	Run: func(cmd *cobra.Command, args []string) {
		// Create new env obj
		env := config.NewEnvConfig()

		// Load env
		if err := env.Load(); err != nil {
			log.Fatal(err)
		}

		// Pass port and start server
		infrastructure.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
