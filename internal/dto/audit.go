package dto

type AuditLogResponse struct {
	ID        uint64 `json:"id"`
	UserID    *uint64 `json:"user_id,omitempty"`
	Action    string  `json:"action"`
	Entity    string  `json:"entity"`
	EntityID  uint64  `json:"entity_id"`
	OldValues *string `json:"old_values,omitempty"`
	NewValues *string `json:"new_values,omitempty"`
	CreatedAt string  `json:"created_at"`
}
