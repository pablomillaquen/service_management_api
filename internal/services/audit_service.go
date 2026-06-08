package services

import (
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
)

type AuditService struct {
	repo *repositories.AuditRepository
}

func NewAuditService(repo *repositories.AuditRepository) *AuditService {
	return &AuditService{repo: repo}
}

func (s *AuditService) FindAll(page, perPage int, filters map[string]interface{}) ([]dto.AuditLogResponse, int64, error) {
	logs, total, err := s.repo.FindAll(page, perPage, filters)
	if err != nil {
		return nil, 0, err
	}
	var responses []dto.AuditLogResponse
	for _, l := range logs {
		ts := l.CreatedAt.Format("2006-01-02T15:04:05Z")
		responses = append(responses, dto.AuditLogResponse{
			ID: l.ID, UserID: l.UserID, Action: string(l.Action),
			Entity: l.Entity, EntityID: l.EntityID,
			OldValues: l.OldValues, NewValues: l.NewValues, CreatedAt: ts,
		})
	}
	return responses, total, nil
}
