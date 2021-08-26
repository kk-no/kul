package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "base",
	Short: "base64 encode/decode tools",
}

var base64EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "base64 encoder",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		str := args[0]
		fmt.Println(encodeBase64(str))
	},
}

var base64DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Base64 decoder",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		str := args[0]
		fmt.Println(decodeBase64(str))
	},
}

func encodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func decodeBase64(str string) string {
	b, _ := base64.StdEncoding.DecodeString(str)
	return string(b)
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.AddCommand(
		base64EncodeCmd,
		base64DecodeCmd,
	)
}
