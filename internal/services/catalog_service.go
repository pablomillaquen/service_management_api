package services

import (
	"fmt"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/equipment"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
)

type EquipmentTypeService struct {
	repo *repositories.EquipmentTypeRepository
}

func NewEquipmentTypeService(repo *repositories.EquipmentTypeRepository) *EquipmentTypeService {
	return &EquipmentTypeService{repo: repo}
}

func (s *EquipmentTypeService) Create(req dto.CreateEquipmentTypeRequest) (*dto.EquipmentTypeResponse, error) {
	et := &equipment.EquipmentType{Name: req.Name}
	if err := s.repo.Create(et); err != nil {
		return nil, fmt.Errorf("equipment type already exists")
	}
	return &dto.EquipmentTypeResponse{ID: et.ID, Name: et.Name}, nil
}

func (s *EquipmentTypeService) FindByID(id uint64) (*dto.EquipmentTypeResponse, error) {
	et, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("equipment type not found")
	}
	return &dto.EquipmentTypeResponse{ID: et.ID, Name: et.Name}, nil
}

func (s *EquipmentTypeService) FindAll() ([]dto.EquipmentTypeResponse, error) {
	types, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var res []dto.EquipmentTypeResponse
	for _, t := range types {
		res = append(res, dto.EquipmentTypeResponse{ID: t.ID, Name: t.Name})
	}
	return res, nil
}

func (s *EquipmentTypeService) Update(id uint64, req dto.UpdateEquipmentTypeRequest) (*dto.EquipmentTypeResponse, error) {
	et, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("equipment type not found")
	}
	et.Name = req.Name
	if err := s.repo.Update(et); err != nil {
		return nil, err
	}
	return &dto.EquipmentTypeResponse{ID: et.ID, Name: et.Name}, nil
}

func (s *EquipmentTypeService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

type BrandService struct {
	repo *repositories.BrandRepository
}

func NewBrandService(repo *repositories.BrandRepository) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) Create(req dto.CreateBrandRequest) (*dto.BrandResponse, error) {
	b := &equipment.Brand{Name: req.Name, EquipmentTypeID: req.EquipmentTypeID}
	if err := s.repo.Create(b); err != nil {
		return nil, err
	}
	return &dto.BrandResponse{ID: b.ID, Name: b.Name, EquipmentTypeID: b.EquipmentTypeID}, nil
}

func (s *BrandService) FindByID(id uint64) (*dto.BrandResponse, error) {
	b, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("brand not found")
	}
	return &dto.BrandResponse{ID: b.ID, Name: b.Name, EquipmentTypeID: b.EquipmentTypeID}, nil
}

func (s *BrandService) FindByTypeID(typeID uint64) ([]dto.BrandResponse, error) {
	brands, err := s.repo.FindByTypeID(typeID)
	if err != nil {
		return nil, err
	}
	var res []dto.BrandResponse
	for _, b := range brands {
		res = append(res, dto.BrandResponse{ID: b.ID, Name: b.Name, EquipmentTypeID: b.EquipmentTypeID})
	}
	return res, nil
}

func (s *BrandService) Update(id uint64, req dto.UpdateBrandRequest) (*dto.BrandResponse, error) {
	b, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("brand not found")
	}
	b.Name = req.Name
	b.EquipmentTypeID = req.EquipmentTypeID
	if err := s.repo.Update(b); err != nil {
		return nil, err
	}
	return &dto.BrandResponse{ID: b.ID, Name: b.Name, EquipmentTypeID: b.EquipmentTypeID}, nil
}

func (s *BrandService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

type EquipmentModelService struct {
	repo *repositories.EquipmentModelRepository
}

func NewEquipmentModelService(repo *repositories.EquipmentModelRepository) *EquipmentModelService {
	return &EquipmentModelService{repo: repo}
}

func (s *EquipmentModelService) Create(req dto.CreateEquipmentModelRequest) (*dto.EquipmentModelResponse, error) {
	m := &equipment.EquipmentModel{
		Name: req.Name, BrandID: req.BrandID, EquipmentTypeID: req.EquipmentTypeID,
	}
	if err := s.repo.Create(m); err != nil {
		return nil, err
	}
	return s.toResponse(m), nil
}

func (s *EquipmentModelService) FindByID(id uint64) (*dto.EquipmentModelResponse, error) {
	m, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("model not found")
	}
	return s.toResponse(m), nil
}

func (s *EquipmentModelService) FindByBrandID(brandID uint64) ([]dto.EquipmentModelResponse, error) {
	models, err := s.repo.FindByBrandID(brandID)
	if err != nil {
		return nil, err
	}
	var res []dto.EquipmentModelResponse
	for _, m := range models {
		res = append(res, *s.toResponse(&m))
	}
	return res, nil
}

func (s *EquipmentModelService) Update(id uint64, req dto.UpdateEquipmentModelRequest) (*dto.EquipmentModelResponse, error) {
	m, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("model not found")
	}
	m.Name = req.Name
	m.BrandID = req.BrandID
	m.EquipmentTypeID = req.EquipmentTypeID
	if err := s.repo.Update(m); err != nil {
		return nil, err
	}
	return s.toResponse(m), nil
}

func (s *EquipmentModelService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *EquipmentModelService) toResponse(m *equipment.EquipmentModel) *dto.EquipmentModelResponse {
	return &dto.EquipmentModelResponse{
		ID: m.ID, Name: m.Name, BrandID: m.BrandID, EquipmentTypeID: m.EquipmentTypeID,
	}
}
