package cmd

import (
	"fmt"

	"github.com/rivo/uniseg"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "count string length",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		str := args[0]
		fmt.Println(countLen(str))
	},
}

func countLen(str string) int {
	gr := uniseg.NewGraphemes(str)

	var c int
	for gr.Next() {
		c++
	}
	return c
}

func init() {
	rootCmd.AddCommand(countCmd)
}
