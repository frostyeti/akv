package cmd

import (
	"fmt"

	"github.com/gobwas/glob"
	"github.com/spf13/cobra"
)

func newCertificateListCmd() *cobra.Command {
	var pattern string

	cmd := &cobra.Command{
		Use:     "ls [pattern]",
		Aliases: []string{"list"},
		Short:   "List certificates",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				pattern = args[0]
			}

			service, err := certificateServiceFactory(cmd)
			if err != nil {
				return err
			}

			certs, err := service.List(cmd.Context())
			if err != nil {
				return err
			}

			if pattern != "" {
				matcher, err := glob.Compile(pattern)
				if err != nil {
					return fmt.Errorf("invalid pattern: %w", err)
				}

				filtered := make([]string, 0, len(certs))
				for _, name := range certs {
					if matcher.Match(name) {
						filtered = append(filtered, name)
					}
				}
				certs = filtered
			}

			for _, name := range certs {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), name)
			}
			return nil
		},
	}

	return cmd
}
