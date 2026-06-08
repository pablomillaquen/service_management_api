package services

import (
	"fmt"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/material"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
)

type MaterialService struct {
	repo *repositories.MaterialRepository
}

func NewMaterialService(repo *repositories.MaterialRepository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (s *MaterialService) Create(req dto.CreateMaterialRequest) (*dto.MaterialResponse, error) {
	existing, _ := s.repo.FindByCode(req.Code)
	if existing != nil {
		return nil, fmt.Errorf("material with this code already exists")
	}
	m := &material.Material{
		Code: req.Code, Description: req.Description, UnitCost: req.UnitCost,
	}
	if err := s.repo.Create(m); err != nil {
		return nil, err
	}
	return s.toResponse(m), nil
}

func (s *MaterialService) FindByID(id uint64) (*dto.MaterialResponse, error) {
	m, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("material not found")
	}
	return s.toResponse(m), nil
}

func (s *MaterialService) FindAll(page, perPage int) ([]dto.MaterialResponse, int64, error) {
	materials, total, err := s.repo.FindAll(page, perPage)
	if err != nil {
		return nil, 0, err
	}
	var responses []dto.MaterialResponse
	for _, m := range materials {
		responses = append(responses, *s.toResponse(&m))
	}
	return responses, total, nil
}

func (s *MaterialService) Update(id uint64, req dto.UpdateMaterialRequest) (*dto.MaterialResponse, error) {
	existing, _ := s.repo.FindByCodeExcludingID(req.Code, id)
	if existing != nil {
		return nil, fmt.Errorf("material with this code already exists")
	}
	m, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	m.Code = req.Code
	m.Description = req.Description
	m.UnitCost = req.UnitCost
	if err := s.repo.Update(m); err != nil {
		return nil, err
	}
	return s.toResponse(m), nil
}

func (s *MaterialService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *MaterialService) toResponse(m *material.Material) *dto.MaterialResponse {
	return &dto.MaterialResponse{
		ID: m.ID, Code: m.Code, Description: m.Description, UnitCost: m.UnitCost,
	}
}
