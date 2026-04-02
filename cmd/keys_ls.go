package cmd

import (
	"fmt"

	"github.com/gobwas/glob"
	"github.com/spf13/cobra"
)

func newKeyListCmd() *cobra.Command {
	var pattern string

	cmd := &cobra.Command{
		Use:     "ls [pattern]",
		Aliases: []string{"list"},
		Short:   "List keys",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				pattern = args[0]
			}

			service, err := keyServiceFactory(cmd)
			if err != nil {
				return err
			}

			keys, err := service.List(cmd.Context())
			if err != nil {
				return err
			}

			if pattern != "" {
				matcher, err := glob.Compile(pattern)
				if err != nil {
					return fmt.Errorf("invalid pattern: %w", err)
				}

				filtered := make([]string, 0, len(keys))
				for _, name := range keys {
					if matcher.Match(name) {
						filtered = append(filtered, name)
					}
				}
				keys = filtered
			}

			for _, name := range keys {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), name)
			}
			return nil
		},
	}

	return cmd
}
