package dto

type CreateEquipmentTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateEquipmentTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

type EquipmentTypeResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type CreateBrandRequest struct {
	Name            string `json:"name" binding:"required"`
	EquipmentTypeID uint64 `json:"equipment_type_id" binding:"required"`
}

type UpdateBrandRequest struct {
	Name            string `json:"name" binding:"required"`
	EquipmentTypeID uint64 `json:"equipment_type_id" binding:"required"`
}

type BrandResponse struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	EquipmentTypeID uint64 `json:"equipment_type_id"`
}

type CreateEquipmentModelRequest struct {
	Name            string `json:"name" binding:"required"`
	BrandID         uint64 `json:"brand_id" binding:"required"`
	EquipmentTypeID uint64 `json:"equipment_type_id" binding:"required"`
}

type UpdateEquipmentModelRequest struct {
	Name            string `json:"name" binding:"required"`
	BrandID         uint64 `json:"brand_id" binding:"required"`
	EquipmentTypeID uint64 `json:"equipment_type_id" binding:"required"`
}

type EquipmentModelResponse struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	BrandID         uint64 `json:"brand_id"`
	EquipmentTypeID uint64 `json:"equipment_type_id"`
}
