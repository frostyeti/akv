package cmd

import "strings"

func toScreamingSnakeCase(input string) string {
	var b strings.Builder
	for i, r := range input {
		switch {
		case r >= 'A' && r <= 'Z':
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(r)
		case r >= 'a' && r <= 'z':
			if i > 0 && input[i-1] >= 'A' && input[i-1] <= 'Z' {
				b.WriteByte('_')
			}
			b.WriteRune(r - ('a' - 'A'))
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		default:
			if i > 0 && input[i-1] != '_' {
				b.WriteByte('_')
			}
		}
	}
	return b.String()
}
