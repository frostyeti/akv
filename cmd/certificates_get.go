package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCertificateGetCmd() *cobra.Command {
	var version string

	cmd := &cobra.Command{
		Use:   "get <name>",
		Short: "Get certificate metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := certificateServiceFactory(cmd)
			if err != nil {
				return err
			}

			info, err := service.Get(cmd.Context(), args[0], version)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(cmd.OutOrStdout(), "id=%s contentType=%s\n", info.ID, info.ContentType)
			return err
		},
	}

	cmd.Flags().StringVar(&version, "version", "", "Specific certificate version")
	return cmd
}
