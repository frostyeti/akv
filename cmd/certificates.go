package cmd

import "github.com/spf13/cobra"

func newCertificatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certificates",
		Short: "Manage Azure Key Vault certificates",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get certificate metadata",
		RunE:  newNotImplementedAction("certificates get"),
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "set",
		Short: "Create or import a certificate",
		RunE:  newNotImplementedAction("certificates set"),
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "rm",
		Short: "Delete a certificate",
		RunE:  newNotImplementedAction("certificates rm"),
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "purge",
		Short: "Permanently purge a deleted certificate",
		RunE:  newNotImplementedAction("certificates purge"),
	})

	return cmd
}
