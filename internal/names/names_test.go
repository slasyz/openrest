package names

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathToPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "simple",
			path:     "/auth/register",
			expected: "AuthRegister",
		},
	}

	for _, tt := range tests {
		result := PathToPascalCase(tt.path)
		assert.Equal(t, tt.expected, result)
	}
}

func TestPathToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "simple",
			path:     "/auth/register",
			expected: "auth_register",
		},
	}

	for _, tt := range tests {
		result := PathToSnakeCase(tt.path)
		assert.Equal(t, tt.expected, result)
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		expected string
	}{
		{
			name:     "first char is not a letter",
			src:      "/auth/register",
			expected: "/auth/register",
		},
		{
			name:     "len = 0",
			src:      "",
			expected: "",
		},
		{
			name:     "len = 1",
			src:      "x",
			expected: "X",
		},
		{
			name:     "len = 7",
			src:      "testBla",
			expected: "TestBla",
		},
	}

	for _, tt := range tests {
		result := Capitalize(tt.src)
		assert.Equal(t, tt.expected, result)
	}
}
