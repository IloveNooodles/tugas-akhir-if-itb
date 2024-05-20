package device

type CreateRequest struct {
	Name     string `json:"name" validate:"required,printascii"`
	Type     string `json:"type" validate:"required,printascii"`
	Labels   string `json:"labels" validate:"required,printascii"`
	NodeName string `json:"node_name" validate:"omitempty,printascii"`
}
