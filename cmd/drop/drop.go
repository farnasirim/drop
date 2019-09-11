package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

	// rootCmd.Flags().String("folder", "",
	// 	"Name of folder in Google Drive that contains your google docs")

	viper.BindPFlags(rootCmd.Flags())
	viper.SetEnvPrefix("DROP")

	serveInit()

	rootCmd.AddCommand(serveCmd)
}

func digestFunc(cmd *cobra.Command, args []string) {
	// fromAddr := viper.GetString("smtp-user")
	// password := viper.GetString("smtp-pass")
}
