package app

import (
	"GoAuth/src/bootstrap"
	"github.com/spf13/cobra"
	"log"
)

// App Commands for interacting with apps
var App = &cobra.Command{
	Use:   "app",
	Short: "Commands for interacting with apps.",
}

func init() {
	App.AddCommand(
		bootstrapCmd,
	)
}

// bootstrapCmd is sub command for App that bootstrap the application
// usage: go run main.go app bootstrap or ./binary app bootstrap
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstraps the application and it's related services.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := bootstrap.Init(); err != nil {
			log.Fatalf("Bootstrap Service: Failed to Initialize. %v", err)
		}

		log.Println("Application exited successfully.")
	},
}
