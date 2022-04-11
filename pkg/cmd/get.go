package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get secret",
	Long:  `Get secret from Secret Service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getSecret(cmd, args)
	},
}

func getSecret(cmd *cobra.Command, args []string) error {
	fmt.Println(cmd)
	fmt.Println(args)

	return nil
}
