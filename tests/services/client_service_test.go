package services

import (
	"testing"

	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientService_Create(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestClientService(db)

	resp, err := svc.Create(dto.CreateClientRequest{
		BusinessName: "Test Corp", TaxID: "12345678901",
		PrimaryContact: "John", Email: "john@test.com",
		Phone: "555-0000", Address: "123 Main St",
	})
	require.NoError(t, err)
	assert.Equal(t, "Test Corp", resp.BusinessName)
	assert.Equal(t, "12345678901", resp.TaxID)
}

func TestClientService_DuplicateTaxID(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestClientService(db)

	svc.Create(dto.CreateClientRequest{
		BusinessName: "Test Corp", TaxID: "12345678901",
		PrimaryContact: "John", Email: "john@test.com",
		Phone: "555-0000", Address: "123 Main St",
	})
	_, err := svc.Create(dto.CreateClientRequest{
		BusinessName: "Other Corp", TaxID: "12345678901",
		PrimaryContact: "Jane", Email: "jane@test.com",
		Phone: "555-0001", Address: "456 Oak St",
	})
	assert.Error(t, err)
}

func TestClientService_FindAll(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestClientService(db)

	svc.Create(dto.CreateClientRequest{
		BusinessName: "Test Corp", TaxID: "11111111111",
		PrimaryContact: "John", Email: "john@test.com",
		Phone: "555-0000", Address: "123 Main St",
	})
	clients, total, err := svc.FindAll(1, 10, "")
	require.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, clients, 1)
}
