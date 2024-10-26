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

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(
		app.App,
	)
}

func initConfig() {
}
