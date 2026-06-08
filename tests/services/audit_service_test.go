package services

import (
	"testing"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/audit"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuditService_FindAll(t *testing.T) {
	db := tests.NewTestDB()
	repo := repositories.NewAuditRepository(db)
	svc := services.NewAuditService(repo)

	uid := uint64(1)
	err := repo.Create(&audit.AuditLog{
		UserID: &uid, Action: audit.ActionInsert,
		Entity: "work_order", EntityID: 1,
	})
	require.NoError(t, err)

	logs, total, err := svc.FindAll(1, 10, map[string]interface{}{})
	require.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, logs, 1)
	assert.Equal(t, "INSERT", logs[0].Action)
}

func TestAuditService_FilterByEntity(t *testing.T) {
	db := tests.NewTestDB()
	repo := repositories.NewAuditRepository(db)
	svc := services.NewAuditService(repo)

	uid := uint64(1)
	require.NoError(t, repo.Create(&audit.AuditLog{
		UserID: &uid, Action: audit.ActionInsert,
		Entity: "work_order", EntityID: 1,
	}))
	require.NoError(t, repo.Create(&audit.AuditLog{
		UserID: &uid, Action: audit.ActionUpdate,
		Entity: "user", EntityID: 2,
	}))

	logs, total, err := svc.FindAll(1, 10, map[string]interface{}{"entity": "entity"})
	require.NoError(t, err)
	assert.Equal(t, int64(0), total)
	assert.Len(t, logs, 0)

	logs, total, err = svc.FindAll(1, 10, map[string]interface{}{"entity": "work_order"})
	require.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, logs, 1)
	assert.Equal(t, "work_order", logs[0].Entity)
}
