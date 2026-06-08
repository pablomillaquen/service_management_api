package dto

type CreateClientRequest struct {
	BusinessName   string `json:"business_name" binding:"required"`
	TaxID          string `json:"tax_id" binding:"required"`
	PrimaryContact string `json:"primary_contact" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required"`
	Address        string `json:"address" binding:"required"`
}

type UpdateClientRequest struct {
	BusinessName   string `json:"business_name" binding:"required"`
	TaxID          string `json:"tax_id" binding:"required"`
	PrimaryContact string `json:"primary_contact" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Phone          string `json:"phone" binding:"required"`
	Address        string `json:"address" binding:"required"`
}

type ClientResponse struct {
	ID             uint64 `json:"id"`
	BusinessName   string `json:"business_name"`
	TaxID          string `json:"tax_id"`
	PrimaryContact string `json:"primary_contact"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
}
