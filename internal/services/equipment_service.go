package services

import (
	"fmt"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/equipment"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
)

type EquipmentService struct {
	repo *repositories.EquipmentRepository
}

func NewEquipmentService(repo *repositories.EquipmentRepository) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (s *EquipmentService) Create(req dto.CreateEquipmentRequest) (*dto.EquipmentResponse, error) {
	existing, _ := s.repo.FindBySerialNumber(req.SerialNumber)
	if existing != nil {
		return nil, fmt.Errorf("equipment with this serial number already exists")
	}
	e := &equipment.Equipment{
		ClientID:     req.ClientID,
		ModelID:      req.ModelID,
		SerialNumber: req.SerialNumber,
		Location:     req.Location,
		Status:       req.Status,
	}
	if err := s.repo.Create(e); err != nil {
		return nil, err
	}
	return s.toResponse(e), nil
}

func (s *EquipmentService) FindByID(id uint64) (*dto.EquipmentResponse, error) {
	e, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("equipment not found")
	}
	return s.toResponse(e), nil
}

func (s *EquipmentService) FindAll(page, perPage int, search string, clientID uint64) ([]dto.EquipmentResponse, int64, error) {
	equipments, total, err := s.repo.FindAll(page, perPage, search, clientID)
	if err != nil {
		return nil, 0, err
	}
	var responses []dto.EquipmentResponse
	for _, e := range equipments {
		responses = append(responses, *s.toResponse(&e))
	}
	return responses, total, nil
}

func (s *EquipmentService) Update(id uint64, req dto.UpdateEquipmentRequest) (*dto.EquipmentResponse, error) {
	existing, _ := s.repo.FindBySerialExcludingID(req.SerialNumber, id)
	if existing != nil {
		return nil, fmt.Errorf("equipment with this serial number already exists")
	}
	e, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	e.ClientID = req.ClientID
	e.ModelID = req.ModelID
	e.SerialNumber = req.SerialNumber
	e.Location = req.Location
	e.Status = req.Status
	if err := s.repo.Update(e); err != nil {
		return nil, err
	}
	return s.toResponse(e), nil
}

func (s *EquipmentService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *EquipmentService) toResponse(e *equipment.Equipment) *dto.EquipmentResponse {
	return &dto.EquipmentResponse{
		ID: e.ID, ClientID: e.ClientID, ModelID: e.ModelID,
		SerialNumber: e.SerialNumber, Location: e.Location, Status: e.Status,
	}
}
