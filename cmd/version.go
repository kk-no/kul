package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const cliVersion = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "CLI Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kul", cliVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
