package cmd

import (
	"fmt"

	"github.com/gobwas/glob"
	"github.com/spf13/cobra"
)

func newSecretListCmd() *cobra.Command {
	var filterPattern string

	cmd := &cobra.Command{
		Use:     "ls [pattern]",
		Short:   "List secrets",
		Aliases: []string{"list"},
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				filterPattern = args[0]
			}

			secrets, err := listSecretNames(cmd)
			if err != nil {
				return err
			}

			if filterPattern != "" {
				matcher, err := glob.Compile(filterPattern)
				if err != nil {
					return fmt.Errorf("invalid pattern: %w", err)
				}

				filtered := make([]string, 0, len(secrets))
				for _, secret := range secrets {
					if matcher.Match(secret) {
						filtered = append(filtered, secret)
					}
				}
				secrets = filtered
			}

			if len(secrets) == 0 {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "No secrets found.")
				return nil
			}

			for _, secret := range secrets {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), secret)
			}

			return nil
		},
	}

	return cmd
}

func listSecretNames(cmd *cobra.Command) ([]string, error) {
	service, err := secretServiceFactory(cmd)
	if err != nil {
		return nil, err
	}

	return service.List(cmd.Context())
}
