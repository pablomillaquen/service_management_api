package services

import (
	"testing"

	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTechnicianLocationService_Create(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestTechnicianLocationService(db)

	resp, err := svc.Create(1, dto.CreateTechnicianLocationRequest{
		Latitude:  -33.4569, Longitude: -70.6483,
	})
	require.NoError(t, err)
	assert.Equal(t, uint64(1), resp.UserID)
	assert.Equal(t, -33.4569, resp.Latitude)
	assert.Equal(t, -70.6483, resp.Longitude)
	assert.NotEmpty(t, resp.CreatedAt)
}

func TestTechnicianLocationService_Create_InvalidLatitude(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestTechnicianLocationService(db)

	_, err := svc.Create(1, dto.CreateTechnicianLocationRequest{
		Latitude: -100, Longitude: 0,
	})
	assert.ErrorIs(t, err, services.ErrInvalidLatitude)
}

func TestTechnicianLocationService_Create_InvalidLongitude(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestTechnicianLocationService(db)

	_, err := svc.Create(1, dto.CreateTechnicianLocationRequest{
		Latitude: 0, Longitude: 200,
	})
	assert.ErrorIs(t, err, services.ErrInvalidLongitude)
}

func TestTechnicianLocationService_FindAll(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestTechnicianLocationService(db)

	_, err := svc.Create(1, dto.CreateTechnicianLocationRequest{Latitude: -33.4569, Longitude: -70.6483})
	require.NoError(t, err)
	_, err = svc.Create(2, dto.CreateTechnicianLocationRequest{Latitude: 40.7128, Longitude: -74.0060})
	require.NoError(t, err)

	locations, total, err := svc.FindAll(1, 10, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, locations, 2)
}

func TestTechnicianLocationService_FindAll_FilterByUser(t *testing.T) {
	db := tests.NewTestDB()
	svc := tests.NewTestTechnicianLocationService(db)

	_, err := svc.Create(1, dto.CreateTechnicianLocationRequest{Latitude: -33.4569, Longitude: -70.6483})
	require.NoError(t, err)
	_, err = svc.Create(2, dto.CreateTechnicianLocationRequest{Latitude: 40.7128, Longitude: -74.0060})
	require.NoError(t, err)

	uid := uint64(1)
	locations, total, err := svc.FindAll(1, 10, &uid, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, locations, 1)
	assert.Equal(t, uint64(1), locations[0].UserID)
}
