package cmd

import "github.com/spf13/cobra"

func newKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Manage Azure Key Vault keys",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get key metadata",
		RunE:  newNotImplementedAction("keys get"),
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "set",
		Short: "Create or update a key",
		RunE:  newNotImplementedAction("keys set"),
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "rm",
		Short: "Delete a key",
		RunE:  newNotImplementedAction("keys rm"),
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "purge",
		Short: "Permanently purge a deleted key",
		RunE:  newNotImplementedAction("keys purge"),
	})

	return cmd
}
