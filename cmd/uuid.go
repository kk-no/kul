package cmd

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate UUID string",
	Run: func(cmd *cobra.Command, args []string) {
		isUpper, _ := cmd.Flags().GetBool("upper")
		uid := generateUUID(isUpper)
		fmt.Println(uid)
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().BoolP("upper", "u", false, "Upper case generate")
}

func generateUUID(isUpper bool) string {
	uid := uuid.New().String()
	if isUpper {
		uid = strings.ToUpper(uid)
	}
	return uid
}
