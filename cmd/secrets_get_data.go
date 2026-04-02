package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func newSecretGetDataCmd() *cobra.Command {
	var version string

	cmd := &cobra.Command{
		Use:   "get-data <name>",
		Short: "Get a secret record as JSON",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := secretServiceFactory(cmd)
			if err != nil {
				return err
			}

			secret, err := service.GetData(cmd.Context(), args[0], version)
			if err != nil {
				return err
			}

			data, err := json.MarshalIndent(secret.Secret, "", "  ")
			if err != nil {
				return fmt.Errorf("marshal secret data: %w", err)
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(data))
			return err
		},
	}

	cmd.Flags().StringVar(&version, "version", "", "Specific secret version GUID")
	return cmd
}
