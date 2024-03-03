package cmd

import (
	"coupon_service/config"
	"log"

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
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
