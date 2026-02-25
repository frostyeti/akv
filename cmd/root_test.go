package cmd

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
)

func TestNewRootCmdContainsExpectedCommands(t *testing.T) {
	root := NewRootCmd()

	expected := []string{"secrets", "keys", "certificates", "get", "set", "rm", "ensure"}
	for _, name := range expected {
		if root.CommandPath() == name {
			t.Fatalf("unexpected command path collision for %q", name)
		}

		if cmd, _, err := root.Find([]string{name}); err != nil || cmd == nil {
			t.Fatalf("expected command %q to exist", name)
		}
	}
}

func TestSecretRootAliasesMatchSecretSubcommands(t *testing.T) {
	tests := []string{"get", "set", "rm", "ensure"}

	for _, action := range tests {
		aliasErr := executeCommand(NewRootCmd(), action)
		secretErr := executeCommand(NewRootCmd(), "secrets", action)

		if !errors.Is(aliasErr, ErrNotImplemented) {
			t.Fatalf("expected alias %q to return ErrNotImplemented, got %v", action, aliasErr)
		}
		if !errors.Is(secretErr, ErrNotImplemented) {
			t.Fatalf("expected secrets %q to return ErrNotImplemented, got %v", action, secretErr)
		}
		if aliasErr.Error() != secretErr.Error() {
			t.Fatalf("expected alias and subcommand to match for %q, got %q and %q", action, aliasErr.Error(), secretErr.Error())
		}
	}
}

func executeCommand(root *cobra.Command, args ...string) error {
	root.SetArgs(args)
	return root.Execute()
}
