package cmd

import (
	"GoAuth/cmd/app"
	"GoAuth/src/config"
	"github.com/spf13/cobra"
	"log"
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
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
}
