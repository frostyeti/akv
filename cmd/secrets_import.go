package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type secretImport struct {
	Value     string            `json:"value"`
	Ensure    bool              `json:"ensure"`
	Size      int               `json:"size"`
	NoUpper   bool              `json:"noUpper"`
	NoLower   bool              `json:"noLower"`
	NoDigits  bool              `json:"noDigits"`
	NoSpecial bool              `json:"noSpecial"`
	Special   string            `json:"special"`
	Chars     string            `json:"chars"`
	Tags      map[string]string `json:"tags"`
	Version   string            `json:"version"`
}

func newSecretsImportCmd() *cobra.Command {
	var (
		file  string
		stdin bool
	)

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import secrets from JSON",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if file == "" && !stdin {
				return fmt.Errorf("must specify either --file or --stdin")
			}
			if file != "" && stdin {
				return fmt.Errorf("--file and --stdin are mutually exclusive")
			}

			payload, err := readJSONInput(file, stdin)
			if err != nil {
				return err
			}

			var raw map[string]json.RawMessage
			if err := json.Unmarshal(payload, &raw); err != nil {
				return fmt.Errorf("parse JSON: %w", err)
			}

			service, err := secretServiceFactory(cmd)
			if err != nil {
				return err
			}

			for name, data := range raw {
				value, err := decodeSecretImport(data)
				if err != nil {
					return fmt.Errorf("parse secret %q: %w", name, err)
				}

				if value.Ensure && value.Value == "" {
					value.Value, err = generateSecretValue(secretGenerationOptions{
						Size:      defaultImportSize(value.Size),
						NoUpper:   value.NoUpper,
						NoLower:   value.NoLower,
						NoDigits:  value.NoDigits,
						NoSpecial: value.NoSpecial,
						Special:   value.Special,
						Chars:     value.Chars,
					})
					if err != nil {
						return err
					}
				}

				if err := service.Set(cmd.Context(), name, value.Value); err != nil {
					return err
				}
			}

			_, err = fmt.Fprintf(cmd.OutOrStdout(), "imported %d secret(s)\n", len(raw))
			return err
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "Input JSON file path")
	cmd.Flags().BoolVar(&stdin, "stdin", false, "Read JSON from stdin")
	return cmd
}

func readJSONInput(file string, stdin bool) ([]byte, error) {
	if file != "" {
		return os.ReadFile(file)
	}

	if !stdin {
		return nil, fmt.Errorf("must specify either --file or --stdin")
	}

	return os.ReadFile(os.Stdin.Name())
}

func decodeSecretImport(data []byte) (secretImport, error) {
	var simple string
	if err := json.Unmarshal(data, &simple); err == nil {
		return secretImport{Value: simple}, nil
	}

	var obj secretImport
	if err := json.Unmarshal(data, &obj); err != nil {
		return secretImport{}, err
	}

	return obj, nil
}

func defaultImportSize(size int) int {
	if size > 0 {
		return size
	}
	return 32
}
