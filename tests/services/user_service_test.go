package services

import (
	"testing"

	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_Create(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestUserService(db)

	resp, err := svc.Create(dto.CreateUserRequest{
		Name: "John", Email: "john@test.com",
		Password: "TestPass1", Role: "technician",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "John", resp.Name)
	assert.Equal(t, "technician", resp.Role)
	assert.True(t, resp.Active)
	assert.NotZero(t, resp.ID)
}

func TestUserService_CreateDuplicateEmail(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestUserService(db)

	_, err := svc.Create(dto.CreateUserRequest{
		Name: "John", Email: "john@test.com",
		Password: "TestPass1", Role: "technician",
	})
	require.NoError(t, err)

	_, err = svc.Create(dto.CreateUserRequest{
		Name: "Jane", Email: "john@test.com",
		Password: "TestPass1", Role: "viewer",
	})
	assert.Error(t, err)
}

func TestUserService_FindByID(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestUserService(db)

	created, err := svc.Create(dto.CreateUserRequest{
		Name: "John", Email: "john@test.com",
		Password: "TestPass1", Role: "technician",
	})
	require.NoError(t, err)
	require.NotNil(t, created)

	found, err := svc.FindByID(created.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, created.Name, found.Name)
}

func TestUserService_Login(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestUserService(db)

	_, err := svc.Create(dto.CreateUserRequest{
		Name: "John", Email: "john@test.com",
		Password: "TestPass1", Role: "technician",
	})
	require.NoError(t, err)

	resp, err := svc.Login(dto.LoginRequest{
		Email: "john@test.com", Password: "TestPass1",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.AccessToken)
	assert.NotEmpty(t, resp.RefreshToken)
}

func TestUserService_LoginInvalid(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestUserService(db)

	_, err := svc.Login(dto.LoginRequest{
		Email: "unknown@test.com", Password: "TestPass1",
	})
	assert.Error(t, err)
}
