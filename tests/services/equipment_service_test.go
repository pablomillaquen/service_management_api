package services

import (
	"testing"

	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEquipmentService_Create(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestEquipmentService(db)

	resp, err := svc.Create(dto.CreateEquipmentRequest{
		ClientID: 1, ModelID: 1,
		SerialNumber: "SN-001", Location: "Warehouse",
		Status: "active",
	})
	require.NoError(t, err)
	assert.Equal(t, "SN-001", resp.SerialNumber)
}

func TestEquipmentService_DuplicateSerial(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestEquipmentService(db)

	svc.Create(dto.CreateEquipmentRequest{
		ClientID: 1, ModelID: 1,
		SerialNumber: "SN-001", Location: "A",
		Status: "active",
	})
	_, err := svc.Create(dto.CreateEquipmentRequest{
		ClientID: 1, ModelID: 1,
		SerialNumber: "SN-001", Location: "B",
		Status: "active",
	})
	assert.Error(t, err)
}
