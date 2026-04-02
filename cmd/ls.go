package cmd

import "github.com/spf13/cobra"

func newLsCmd() *cobra.Command {
	cmd := newSecretListCmd()
	cmd.Short = "List secrets (alias for 'akv secrets ls')"
	return cmd
}
