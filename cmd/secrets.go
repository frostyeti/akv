package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/frostyeti/akv/internal/config"
	"github.com/frostyeti/akv/internal/keyvault"
	"github.com/spf13/cobra"
)

type secretService interface {
	Get(ctx context.Context, name string, version string) (string, error)
	GetData(ctx context.Context, name string, version string) (keyvault.SecretInfo, error)
	Set(ctx context.Context, name string, value string) error
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, name string, in keyvault.SecretUpdateInput) error
	List(ctx context.Context) ([]string, error)
	Purge(ctx context.Context, name string) error
}

var secretServiceFactory = buildSecretService

func newSecretsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secrets",
		Short: "Manage Azure Key Vault secrets",
	}

	cmd.AddCommand(newSecretGetCmd())
	cmd.AddCommand(newSecretGetDataCmd())
	cmd.AddCommand(newSecretSetCmd())
	cmd.AddCommand(newSecretDeleteCmd())
	cmd.AddCommand(newSecretPurgeCmd())
	cmd.AddCommand(newSecretEnsureCmd())
	cmd.AddCommand(newSecretUpdateCmd())
	cmd.AddCommand(newSecretListCmd())
	cmd.AddCommand(newSecretsImportCmd())
	cmd.AddCommand(newSecretsExportCmd())
	cmd.AddCommand(newSecretsSyncCmd())

	return cmd
}

func newSecretRootAliasCmds() []*cobra.Command {
	return []*cobra.Command{
		newSecretGetAliasCmd(),
		newSecretSetAliasCmd(),
		newSecretDeleteAliasCmd(),
		newSecretEnsureAliasCmd(),
	}
}

func buildSecretService(cmd *cobra.Command) (secretService, error) {
	vaultURL, err := resolveVaultURL(cmd)
	if err != nil {
		return nil, err
	}

	service, err := keyvault.NewSecretsService(vaultURL)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func resolveVaultURL(cmd *cobra.Command) (string, error) {
	// Check command line flag first
	vaultURL, err := cmd.Root().PersistentFlags().GetString("vault-url")
	if err != nil {
		return "", err
	}

	// Then check environment variable
	if vaultURL == "" {
		vaultURL = os.Getenv("AKV_VAULT_URL")
	}

	// Finally check config for current vault
	if vaultURL == "" {
		mgr, err := config.NewManager()
		if err == nil {
			url, err := mgr.GetVaultURL("")
			if err == nil {
				vaultURL = url
			}
		}
	}

	if vaultURL == "" {
		return "", ErrVaultURLRequired
	}

	return vaultURL, nil
}

func handleSecretNotFound(err error) bool {
	return errors.Is(err, keyvault.ErrSecretNotFound)
}
