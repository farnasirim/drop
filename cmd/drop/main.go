package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "drop sub-command [args]",
	Short: "Quick and easy link drop service",
	Long: `Quick and easy link drop service
Almost all of command line flags can also be passed in as
environement variables with 'DROP_' prefix: e.g. --some-arg -> DIG_SOME_ARG
	`,
}
