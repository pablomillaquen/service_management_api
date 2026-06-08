package services

import (
	"fmt"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/client"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
	"gorm.io/gorm"
)

type ClientService struct {
	repo *repositories.ClientRepository
}

func NewClientService(repo *repositories.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Create(req dto.CreateClientRequest) (*dto.ClientResponse, error) {
	existing, _ := s.repo.FindByTaxID(req.TaxID)
	if existing != nil {
		return nil, fmt.Errorf("a client with this tax ID already exists")
	}
	c := &client.Client{
		BusinessName:   req.BusinessName,
		TaxID:          req.TaxID,
		PrimaryContact: req.PrimaryContact,
		Email:          req.Email,
		Phone:          req.Phone,
		Address:        req.Address,
	}
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return s.toResponse(c), nil
}

func (s *ClientService) FindByID(id uint64) (*dto.ClientResponse, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("client not found")
		}
		return nil, err
	}
	return s.toResponse(c), nil
}

func (s *ClientService) FindAll(page, perPage int, search string) ([]dto.ClientResponse, int64, error) {
	clients, total, err := s.repo.FindAll(page, perPage, search)
	if err != nil {
		return nil, 0, err
	}
	var responses []dto.ClientResponse
	for _, c := range clients {
		responses = append(responses, *s.toResponse(&c))
	}
	return responses, total, nil
}

func (s *ClientService) Update(id uint64, req dto.UpdateClientRequest) (*dto.ClientResponse, error) {
	existing, _ := s.repo.FindByTaxIDExcludingID(req.TaxID, id)
	if existing != nil {
		return nil, fmt.Errorf("a client with this tax ID already exists")
	}
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	c.BusinessName = req.BusinessName
	c.TaxID = req.TaxID
	c.PrimaryContact = req.PrimaryContact
	c.Email = req.Email
	c.Phone = req.Phone
	c.Address = req.Address
	if err := s.repo.Update(c); err != nil {
		return nil, err
	}
	return s.toResponse(c), nil
}

func (s *ClientService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *ClientService) toResponse(c *client.Client) *dto.ClientResponse {
	return &dto.ClientResponse{
		ID:             c.ID,
		BusinessName:   c.BusinessName,
		TaxID:          c.TaxID,
		PrimaryContact: c.PrimaryContact,
		Email:          c.Email,
		Phone:          c.Phone,
		Address:        c.Address,
	}
}
