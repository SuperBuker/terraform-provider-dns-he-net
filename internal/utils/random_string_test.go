package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
)

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "zero length",
			length: 0,
		},
		{
			name:   "single character",
			length: 1,
		},
		{
			name:   "small string",
			length: 10,
		},
		{
			name:   "medium string",
			length: 50,
		},
		{
			name:   "large string",
			length: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GenerateRandomString(tt.length)

			// Check length
			if len(result) != tt.length {
				t.Errorf("GenerateRandomString(%d) length = %d, want %d", tt.length, len(result), tt.length)
			}

			// Check that all characters are from the valid set
			validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01256789"
			for i, char := range result {
				found := false
				for _, validChar := range validChars {
					if char == validChar {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("GenerateRandomString(%d) contains invalid character '%c' at position %d", tt.length, char, i)
				}
			}
		})
	}

	// Test uniqueness - generate multiple strings and verify they're different
	t.Run("uniqueness", func(t *testing.T) {
		length := 20
		iterations := 100
		seen := make(map[string]bool)

		for i := 0; i < iterations; i++ {
			result := utils.GenerateRandomString(length)
			if seen[result] {
				t.Logf("Warning: duplicate string generated after %d iterations", i)
			}
			seen[result] = true
		}

		// With 20 characters, we should see many unique strings
		if len(seen) < iterations/2 {
			t.Errorf("Generated only %d unique strings out of %d attempts", len(seen), iterations)
		}
	})
}
