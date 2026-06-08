package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("TestPass1")
	require.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.True(t, CheckPassword("TestPass1", hash))
	assert.False(t, CheckPassword("WrongPass1", hash))
}

func TestValidatePasswordPolicy(t *testing.T) {
	tests := []struct {
		name  string
		pwd   string
		valid bool
	}{
		{"valid password", "TestPass1", true},
		{"too short", "Te1", false},
		{"no uppercase", "testpass1", false},
		{"no lowercase", "TESTPASS1", false},
		{"no number", "TestPass", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePasswordPolicy(tt.pwd)
			if tt.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
