package cmd

import (
	"GoAuth/cmd/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-auth",
	Short: "go auth CLI",
	Long:  "Golang Authentication CLI",
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(
		app.App,
	)
}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	// put config needed fore cli command
}
