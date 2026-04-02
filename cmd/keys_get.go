package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newKeyGetCmd() *cobra.Command {
	var version string

	cmd := &cobra.Command{
		Use:   "get <name>",
		Short: "Get key metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := keyServiceFactory(cmd)
			if err != nil {
				return err
			}

			info, err := service.Get(cmd.Context(), args[0], version)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(cmd.OutOrStdout(), "id=%s type=%s\n", info.ID, info.Type)
			return err
		},
	}

	cmd.Flags().StringVar(&version, "version", "", "Specific key version")
	return cmd
}
