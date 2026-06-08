package dto

type CreateMaterialRequest struct {
	Code        string  `json:"code" binding:"required"`
	Description string  `json:"description" binding:"required"`
	UnitCost    float64 `json:"unit_cost" binding:"required"`
}

type UpdateMaterialRequest struct {
	Code        string  `json:"code" binding:"required"`
	Description string  `json:"description" binding:"required"`
	UnitCost    float64 `json:"unit_cost" binding:"required"`
}

type MaterialResponse struct {
	ID          uint64  `json:"id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	UnitCost    float64 `json:"unit_cost"`
}
