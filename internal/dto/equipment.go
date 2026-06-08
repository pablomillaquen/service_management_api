package dto

type CreateEquipmentRequest struct {
	ClientID     uint64 `json:"client_id" binding:"required"`
	ModelID      uint64 `json:"model_id" binding:"required"`
	SerialNumber string `json:"serial_number" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Status       string `json:"status" binding:"required"`
}

type UpdateEquipmentRequest struct {
	ClientID     uint64 `json:"client_id" binding:"required"`
	ModelID      uint64 `json:"model_id" binding:"required"`
	SerialNumber string `json:"serial_number" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Status       string `json:"status" binding:"required"`
}

type EquipmentResponse struct {
	ID           uint64 `json:"id"`
	ClientID     uint64 `json:"client_id"`
	ModelID      uint64 `json:"model_id"`
	SerialNumber string `json:"serial_number"`
	Location     string `json:"location"`
	Status       string `json:"status"`
}
