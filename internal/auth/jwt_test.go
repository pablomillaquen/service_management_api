package auth

import (
	"testing"
	"time"

	"github.com/pablomillaquen/speckit_golang_api/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestJWTService() *JWTService {
	cfg := configs.JWTConfig{
		Secret:            "test-secret-key-12345",
		AccessExpiration:  time.Minute * 15,
		RefreshExpiration: time.Hour * 24,
	}
	return NewJWTService(cfg)
}

func TestGenerateTokenPair(t *testing.T) {
	s := newTestJWTService()
	pair, err := s.GenerateTokenPair(1, "test@test.com", "administrator")
	require.NoError(t, err)
	assert.NotEmpty(t, pair.AccessToken)
	assert.NotEmpty(t, pair.RefreshToken)
	assert.Equal(t, 900, pair.ExpiresIn)
}

func TestValidateToken(t *testing.T) {
	s := newTestJWTService()
	pair, err := s.GenerateTokenPair(1, "test@test.com", "technician")
	require.NoError(t, err)

	claims, err := s.ValidateToken(pair.AccessToken)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), claims.UserID)
	assert.Equal(t, "test@test.com", claims.Email)
	assert.Equal(t, "technician", claims.Role)
}

func TestValidateRefreshToken(t *testing.T) {
	s := newTestJWTService()
	pair, err := s.GenerateTokenPair(1, "test@test.com", "viewer")
	require.NoError(t, err)

	claims, err := s.ValidateToken(pair.RefreshToken)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), claims.UserID)
	assert.Equal(t, "viewer", claims.Role)
}

func TestInvalidToken(t *testing.T) {
	s := newTestJWTService()
	_, err := s.ValidateToken("invalid-token-string")
	assert.Error(t, err)
}
