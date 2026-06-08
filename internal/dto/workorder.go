package dto

type CreateWorkOrderRequest struct {
	ClientID      uint64 `json:"client_id" binding:"required"`
	EquipmentID   uint64 `json:"equipment_id" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Priority      string `json:"priority" binding:"required"`
	ScheduledDate string `json:"scheduled_date" binding:"required"`
}

type UpdateWorkOrderRequest struct {
	ClientID      uint64 `json:"client_id" binding:"required"`
	EquipmentID   uint64 `json:"equipment_id" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Priority      string `json:"priority" binding:"required"`
	ScheduledDate string `json:"scheduled_date" binding:"required"`
}

type AssignTechnicianRequest struct {
	TechnicianID uint64 `json:"technician_id" binding:"required"`
}

type ChangeStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type AddNoteRequest struct {
	Text string `json:"text" binding:"required"`
}

type AddMaterialRequest struct {
	MaterialID uint64  `json:"material_id" binding:"required"`
	Quantity   float64 `json:"quantity" binding:"required"`
}

type WorkOrderResponse struct {
	ID            uint64  `json:"id"`
	ClientID      uint64  `json:"client_id"`
	EquipmentID   uint64  `json:"equipment_id"`
	Description   string  `json:"description"`
	Priority      string  `json:"priority"`
	Status        string  `json:"status"`
	ScheduledDate string  `json:"scheduled_date"`
}

type WorkOrderDetailResponse struct {
	ID            uint64                    `json:"id"`
	ClientID      uint64                    `json:"client_id"`
	EquipmentID   uint64                    `json:"equipment_id"`
	Description   string                    `json:"description"`
	Priority      string                    `json:"priority"`
	Status        string                    `json:"status"`
	ScheduledDate string                    `json:"scheduled_date"`
	Notes         []WorkOrderNoteResponse   `json:"notes,omitempty"`
	Materials     []WorkOrderMaterialResponse `json:"materials,omitempty"`
}

type WorkOrderNoteResponse struct {
	ID       uint64 `json:"id"`
	AuthorID uint64 `json:"author_id"`
	Text     string `json:"text"`
}

type WorkOrderMaterialResponse struct {
	ID         uint64  `json:"id"`
	MaterialID uint64  `json:"material_id"`
	Quantity   float64 `json:"quantity"`
	UserID     uint64  `json:"user_id"`
}
