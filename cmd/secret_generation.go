package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

const defaultSecretSpecialChars = "@_-{}|#!~:^"

type secretGenerationOptions struct {
	Generate  bool
	Size      int
	NoUpper   bool
	NoLower   bool
	NoDigits  bool
	NoSpecial bool
	Special   string
	Chars     string
}

func registerSecretGenerationFlags(cmd *cobra.Command, opts *secretGenerationOptions) {
	cmd.Flags().BoolVar(&opts.Generate, "generate", false, "Generate a random secret value")
	cmd.Flags().IntVar(&opts.Size, "size", 16, "Size of the generated secret in characters")
	cmd.Flags().BoolVarP(&opts.NoUpper, "no-upper", "U", false, "Exclude uppercase letters from generated secret")
	cmd.Flags().BoolVarP(&opts.NoLower, "no-lower", "L", false, "Exclude lowercase letters from generated secret")
	cmd.Flags().BoolVarP(&opts.NoDigits, "no-digits", "D", false, "Exclude digits from generated secret")
	cmd.Flags().BoolVarP(&opts.NoSpecial, "no-special", "S", false, "Exclude special characters from generated secret")
	cmd.Flags().StringVar(&opts.Special, "special", "", "Custom special characters to use")
	cmd.Flags().StringVar(&opts.Chars, "chars", "", "Use only these specific characters (overrides all other character options)")
}

func generateSecretValue(opts secretGenerationOptions) (string, error) {
	if opts.Size <= 0 {
		return "", fmt.Errorf("size must be greater than zero")
	}

	charset := opts.Chars
	if charset == "" {
		if !opts.NoUpper {
			charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
		if !opts.NoLower {
			charset += "abcdefghijklmnopqrstuvwxyz"
		}
		if !opts.NoDigits {
			charset += "0123456789"
		}
		if !opts.NoSpecial {
			if opts.Special != "" {
				charset += opts.Special
			} else {
				charset += defaultSecretSpecialChars
			}
		}
	}

	if charset == "" {
		return "", fmt.Errorf("no characters available for secret generation")
	}

	secret := make([]byte, opts.Size)
	limit := big.NewInt(int64(len(charset)))
	for i := range secret {
		n, err := rand.Int(rand.Reader, limit)
		if err != nil {
			return "", fmt.Errorf("generate random secret: %w", err)
		}
		secret[i] = charset[n.Int64()]
	}

	return string(secret), nil
}
