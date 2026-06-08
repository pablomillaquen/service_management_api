package services

import (
	"testing"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/workorder"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkOrderService_Create(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestWorkOrderService(db)

	resp, err := svc.Create(dto.CreateWorkOrderRequest{
		ClientID: 1, EquipmentID: 1,
		Description: "Fix printer", Priority: "high",
		ScheduledDate: "2026-06-10",
	}, 1)
	require.NoError(t, err)
	assert.Equal(t, string(workorder.StatusPending), resp.Status)
}

func TestWorkOrderService_AssignTechnician(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestWorkOrderService(db)

	resp, _ := svc.Create(dto.CreateWorkOrderRequest{
		ClientID: 1, EquipmentID: 1,
		Description: "Fix printer", Priority: "high",
		ScheduledDate: "2026-06-10",
	}, 1)
	err := svc.AssignTechnician(resp.ID, 2, 1)
	require.NoError(t, err)
}

func TestWorkOrderService_StatusTransition(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestWorkOrderService(db)

	resp, _ := svc.Create(dto.CreateWorkOrderRequest{
		ClientID: 1, EquipmentID: 1,
		Description: "Fix printer", Priority: "high",
		ScheduledDate: "2026-06-10",
	}, 1)

	svc.AssignTechnician(resp.ID, 2, 1)
	err := svc.ChangeStatus(resp.ID, workorder.StatusInProgress, 2)
	require.NoError(t, err)

	err = svc.ChangeStatus(resp.ID, workorder.StatusCompleted, 2)
	require.NoError(t, err)
}

func TestWorkOrderService_InvalidTransition(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestWorkOrderService(db)

	resp, _ := svc.Create(dto.CreateWorkOrderRequest{
		ClientID: 1, EquipmentID: 1,
		Description: "Fix printer", Priority: "high",
		ScheduledDate: "2026-06-10",
	}, 1)
	err := svc.ChangeStatus(resp.ID, workorder.StatusCompleted, 1)
	assert.Error(t, err)
}

func TestWorkOrderService_AddNote(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestWorkOrderService(db)

	resp, _ := svc.Create(dto.CreateWorkOrderRequest{
		ClientID: 1, EquipmentID: 1,
		Description: "Fix printer", Priority: "high",
		ScheduledDate: "2026-06-10",
	}, 1)
	err := svc.AddNote(resp.ID, 1, "Test observation")
	require.NoError(t, err)

	detail, _ := svc.FindByID(resp.ID)
	assert.Len(t, detail.Notes, 1)
	assert.Equal(t, "Test observation", detail.Notes[0].Text)
}

func TestWorkOrderService_AddMaterial(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestWorkOrderService(db)

	resp, _ := svc.Create(dto.CreateWorkOrderRequest{
		ClientID: 1, EquipmentID: 1,
		Description: "Fix printer", Priority: "high",
		ScheduledDate: "2026-06-10",
	}, 1)
	err := svc.AddMaterial(resp.ID, 1, 1, 2.5)
	require.NoError(t, err)

	detail, _ := svc.FindByID(resp.ID)
	assert.Len(t, detail.Materials, 1)
	assert.Equal(t, 2.5, detail.Materials[0].Quantity)
}
