package cmd

import "github.com/spf13/cobra"

func newSecretsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secrets",
		Short: "Manage Azure Key Vault secrets",
	}

	for _, secretCmd := range newSecretOperationCmds() {
		cmd.AddCommand(secretCmd)
	}

	return cmd
}

func newSecretOperationCmds() []*cobra.Command {
	return []*cobra.Command{
		newSecretOperationCmd("get", "Get a secret value"),
		newSecretOperationCmd("set", "Set a secret value"),
		newSecretOperationCmd("rm", "Delete a secret"),
		newSecretOperationCmd("ensure", "Ensure a secret exists"),
	}
}

func newSecretOperationCmd(use string, short string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		RunE:  newNotImplementedAction("secrets " + use),
	}
}

func newSecretRootAliasCmds() []*cobra.Command {
	return []*cobra.Command{
		newSecretRootAliasCmd("get", "Alias for secrets get"),
		newSecretRootAliasCmd("set", "Alias for secrets set"),
		newSecretRootAliasCmd("rm", "Alias for secrets rm"),
		newSecretRootAliasCmd("ensure", "Alias for secrets ensure"),
	}
}

func newSecretRootAliasCmd(use string, short string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		RunE:  newNotImplementedAction("secrets " + use),
	}
}
