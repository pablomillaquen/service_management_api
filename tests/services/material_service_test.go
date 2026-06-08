package services

import (
	"testing"

	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaterialService_Create(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestMaterialService(db)

	resp, err := svc.Create(dto.CreateMaterialRequest{
		Code: "TON-001", Description: "Toner HP", UnitCost: 45.50,
	})
	require.NoError(t, err)
	assert.Equal(t, "TON-001", resp.Code)
	assert.Equal(t, 45.50, resp.UnitCost)
}

func TestMaterialService_DuplicateCode(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestMaterialService(db)

	svc.Create(dto.CreateMaterialRequest{
		Code: "TON-001", Description: "Toner HP", UnitCost: 45.50,
	})
	_, err := svc.Create(dto.CreateMaterialRequest{
		Code: "TON-001", Description: "Toner Dup", UnitCost: 50.00,
	})
	assert.Error(t, err)
}
