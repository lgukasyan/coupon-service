package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start coupon service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("Running...")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
