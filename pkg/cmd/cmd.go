package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// rootCmd is cobra root object
var rootCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "keesecret",
	Short: "keesecret is a exporter for secrets from Secret Service",
	Long: `keesecret uses Secret Service to get secrets and exports
			them as needed env variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
