package services

import (
	"errors"
	"time"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/technicianlocation"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
)

var (
	ErrInvalidLatitude  = errors.New("latitude must be between -90 and 90")
	ErrInvalidLongitude = errors.New("longitude must be between -180 and 180")
)

type TechnicianLocationService struct {
	repo *repositories.TechnicianLocationRepository
}

func NewTechnicianLocationService(repo *repositories.TechnicianLocationRepository) *TechnicianLocationService {
	return &TechnicianLocationService{repo: repo}
}

func (s *TechnicianLocationService) Create(userID uint64, req dto.CreateTechnicianLocationRequest) (*dto.TechnicianLocationResponse, error) {
	if req.Latitude < -90 || req.Latitude > 90 {
		return nil, ErrInvalidLatitude
	}
	if req.Longitude < -180 || req.Longitude > 180 {
		return nil, ErrInvalidLongitude
	}
	loc := &technicianlocation.TechnicianLocation{
		UserID:    userID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	if err := s.repo.Create(loc); err != nil {
		return nil, err
	}
	return &dto.TechnicianLocationResponse{
		ID:        loc.ID,
		UserID:    loc.UserID,
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		CreatedAt: loc.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *TechnicianLocationService) FindAll(page, perPage int, userID *uint64, startDate, endDate *time.Time) ([]dto.TechnicianLocationResponse, int64, error) {
	locations, total, err := s.repo.FindAll(page, perPage, userID, startDate, endDate)
	if err != nil {
		return nil, 0, err
	}
	var responses []dto.TechnicianLocationResponse
	for _, l := range locations {
		responses = append(responses, dto.TechnicianLocationResponse{
			ID:        l.ID,
			UserID:    l.UserID,
			Latitude:  l.Latitude,
			Longitude: l.Longitude,
			CreatedAt: l.CreatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}
	return responses, total, nil
}
