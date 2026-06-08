package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pablomillaquen/speckit_golang_api/internal/domain/audit"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/workorder"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
	"github.com/pablomillaquen/speckit_golang_api/pkg/logger"
)

type WorkOrderService struct {
	woRepo     *repositories.WorkOrderRepository
	noteRepo   *repositories.WorkOrderNoteRepository
	matRepo    *repositories.WorkOrderMaterialRepository
	auditRepo  *repositories.AuditRepository
}

func NewWorkOrderService(
	woRepo *repositories.WorkOrderRepository,
	noteRepo *repositories.WorkOrderNoteRepository,
	matRepo *repositories.WorkOrderMaterialRepository,
	auditRepo *repositories.AuditRepository,
) *WorkOrderService {
	return &WorkOrderService{
		woRepo: woRepo, noteRepo: noteRepo, matRepo: matRepo, auditRepo: auditRepo,
	}
}

func (s *WorkOrderService) Create(req dto.CreateWorkOrderRequest, userID uint64) (*dto.WorkOrderResponse, error) {
	if !workorder.Priority(req.Priority).IsValid() {
		return nil, fmt.Errorf("invalid priority: must be low, medium, high, or critical")
	}
	wo := &workorder.WorkOrder{
		ClientID:      req.ClientID,
		EquipmentID:   req.EquipmentID,
		Description:   req.Description,
		Priority:      workorder.Priority(req.Priority),
		Status:        workorder.StatusPending,
		ScheduledDate: req.ScheduledDate,
	}
	if err := s.woRepo.Create(wo); err != nil {
		return nil, err
	}
	s.logAudit(userID, audit.ActionInsert, "work_order", wo.ID, nil, wo)
	return s.toResponse(wo), nil
}

func (s *WorkOrderService) FindByID(id uint64) (*dto.WorkOrderDetailResponse, error) {
	wo, err := s.woRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("work order not found")
	}
	notes, _ := s.noteRepo.FindByWorkOrderID(id)
	materials, _ := s.matRepo.FindByWorkOrderID(id)
	resp := &dto.WorkOrderDetailResponse{
		ID:            wo.ID,
		ClientID:      wo.ClientID,
		EquipmentID:   wo.EquipmentID,
		Description:   wo.Description,
		Priority:      string(wo.Priority),
		Status:        string(wo.Status),
		ScheduledDate: wo.ScheduledDate,
	}
	for _, n := range notes {
		resp.Notes = append(resp.Notes, dto.WorkOrderNoteResponse{
			ID: n.ID, AuthorID: n.AuthorID, Text: n.Text,
		})
	}
	for _, m := range materials {
		resp.Materials = append(resp.Materials, dto.WorkOrderMaterialResponse{
			ID: m.ID, MaterialID: m.MaterialID, Quantity: m.Quantity, UserID: m.UserID,
		})
	}
	return resp, nil
}

func (s *WorkOrderService) FindAll(page, perPage int, filters map[string]interface{}) ([]dto.WorkOrderResponse, int64, error) {
	orders, total, err := s.woRepo.FindAll(page, perPage, filters)
	if err != nil {
		return nil, 0, err
	}
	var responses []dto.WorkOrderResponse
	for _, wo := range orders {
		responses = append(responses, *s.toResponse(&wo))
	}
	return responses, total, nil
}

func (s *WorkOrderService) Update(id uint64, req dto.UpdateWorkOrderRequest, userID uint64) (*dto.WorkOrderResponse, error) {
	wo, err := s.woRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	old := *wo
	wo.ClientID = req.ClientID
	wo.EquipmentID = req.EquipmentID
	wo.Description = req.Description
	wo.Priority = workorder.Priority(req.Priority)
	wo.ScheduledDate = req.ScheduledDate
	if err := s.woRepo.Update(wo); err != nil {
		return nil, err
	}
	s.logAudit(userID, audit.ActionUpdate, "work_order", wo.ID, &old, wo)
	return s.toResponse(wo), nil
}

func (s *WorkOrderService) Delete(id uint64, userID uint64) error {
	wo, err := s.woRepo.FindByID(id)
	if err != nil {
		return err
	}
	old := *wo
	if err := s.woRepo.Delete(id); err != nil {
		return err
	}
	s.logAudit(userID, audit.ActionDelete, "work_order", id, &old, nil)
	return nil
}

func (s *WorkOrderService) AssignTechnician(id, technicianID, assignedByID uint64) error {
	wo, err := s.woRepo.FindByID(id)
	if err != nil {
		return err
	}
	if wo.Status != workorder.StatusPending {
		return fmt.Errorf("can only assign pending work orders")
	}
	old := *wo
	now := time.Now()
	wo.TechnicianID = &technicianID
	wo.AssignedByID = &assignedByID
	wo.AssignedAt = &now
	wo.Status = workorder.StatusAssigned
	if err := s.woRepo.Update(wo); err != nil {
		return err
	}
	s.logAudit(assignedByID, audit.ActionAssignment, "work_order", id, &old, wo)
	logger.Info("Work order %d assigned to technician %d", id, technicianID)
	return nil
}

func (s *WorkOrderService) ChangeStatus(id uint64, newStatus workorder.Status, userID uint64) error {
	if !newStatus.IsValid() {
		return fmt.Errorf("invalid status: %s", newStatus)
	}
	wo, err := s.woRepo.FindByID(id)
	if err != nil {
		return err
	}
	if wo.Status == workorder.StatusCompleted || wo.Status == workorder.StatusCancelled {
		return fmt.Errorf("cannot change status of a %s work order", wo.Status)
	}
	if !workorder.IsValidTransition(wo.Status, newStatus) {
		return fmt.Errorf("cannot transition from %s to %s", wo.Status, newStatus)
	}
	old := *wo
	wo.Status = newStatus
	if newStatus == workorder.StatusCompleted {
		now := time.Now()
		wo.CompletedDate = &now
	}
	if err := s.woRepo.Update(wo); err != nil {
		return err
	}
	s.logAudit(userID, audit.ActionStatusChange, "work_order", id, &old, wo)
	logger.Info("Work order %d status changed to %s", id, newStatus)
	return nil
}

func (s *WorkOrderService) AddNote(workOrderID, authorID uint64, text string) error {
	note := &workorder.WorkOrderNote{
		WorkOrderID: workOrderID, AuthorID: authorID, Text: text,
	}
	return s.noteRepo.Create(note)
}

func (s *WorkOrderService) AddMaterial(workOrderID, materialID, userID uint64, quantity float64) error {
	wm := &workorder.WorkOrderMaterial{
		WorkOrderID: workOrderID, MaterialID: materialID,
		Quantity: quantity, UserID: userID,
	}
	return s.matRepo.Create(wm)
}

func (s *WorkOrderService) toResponse(wo *workorder.WorkOrder) *dto.WorkOrderResponse {
	return &dto.WorkOrderResponse{
		ID: wo.ID, ClientID: wo.ClientID, EquipmentID: wo.EquipmentID,
		Description: wo.Description, Priority: string(wo.Priority),
		Status: string(wo.Status), ScheduledDate: wo.ScheduledDate,
	}
}

func (s *WorkOrderService) logAudit(userID uint64, action audit.Action, entity string, entityID uint64, old, new interface{}) {
	var oldJSON, newJSON *string
	if old != nil {
		b, _ := json.Marshal(old)
		s := string(b)
		oldJSON = &s
	}
	if new != nil {
		b, _ := json.Marshal(new)
		s := string(b)
		newJSON = &s
	}
	uid := userID
	s.auditRepo.Create(&audit.AuditLog{
		UserID: &uid, Action: action, Entity: entity,
		EntityID: entityID, OldValues: oldJSON, NewValues: newJSON,
	})
}
