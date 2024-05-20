package device

type CreateRequest struct {
	Name     string `json:"name" validate:"required,printascii,min=8"`
	Type     string `json:"type" validate:"required,printascii"`
	Labels   string `json:"labels" validate:"required,printascii,contains=="`
	NodeName string `json:"node_name" validate:"omitempty,printascii,min=8"`
}
