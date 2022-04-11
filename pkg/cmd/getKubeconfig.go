package cmd

import (
	"fmt"
	"os"

	"github.com/einfachnuralex/keesecret/pkg/secretservice"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

//nolint:gochecknoinits
func init() {
	getCmd.AddCommand(getKc)
}

var getKc = &cobra.Command{ //nolint:gochecknoglobals
	Use:     "kubeconfig",
	Aliases: []string{"kc"},
	Short:   "Get kubeconfig",
	Long:    `Get kubeconfig secret from Secret Service`,
	RunE:    getKubeconfig,
}

func getKubeconfig(cmd *cobra.Command, args []string) error {
	secrets, err := secretservice.GetSecrets("kubeconfig")
	if err != nil {
		return err
	}
	idx, err := fuzzyfinder.Find(
		secrets,
		func(i int) string {
			return secrets[i].Title
		})
	if err != nil {
		return err
	}

	// fmt.Printf("selected: %v\n", secrets[idx].Title)
	var filename = "/tmp/kc"
	err = writeFile(filename, secrets[idx].Secret)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(os.Stdout, "KUBECONFIG=%s\n", filename)
	return nil
}

func writeFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
