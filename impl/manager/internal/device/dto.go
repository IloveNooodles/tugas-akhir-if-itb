package device

type CreateRequest struct {
	Name       string `json:"name" validate:"required,printascii"`
	Type       string `json:"type" validate:"required,printascii"`
	Attributes string `json:"attributes" validate:"required,printascii"`
}
