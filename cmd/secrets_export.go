package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newSecretsExportCmd() *cobra.Command {
	var file string

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export secrets to JSON",
		Long: `Export secrets to JSON.

Output shape:
  {
    "db-password": {
      "value": "secret-value",
      "contentType": "text/plain",
      "tags": {"team": "platform"}
    }
  }`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := secretServiceFactory(cmd)
			if err != nil {
				return err
			}

			secrets, err := service.List(cmd.Context())
			if err != nil {
				return err
			}

			output := make(map[string]any, len(secrets))
			for _, name := range secrets {
				secret, err := service.GetData(cmd.Context(), name, "")
				if err != nil {
					return err
				}
				output[name] = secret.Secret
			}

			data, err := json.MarshalIndent(output, "", "  ")
			if err != nil {
				return fmt.Errorf("marshal export: %w", err)
			}

			if file != "" {
				return os.WriteFile(file, data, 0600)
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(data))
			return err
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "Output JSON file path (JSON format)")
	return cmd
}
