package groups

type CreateRequest struct {
	Name string `json:"name" validate:"required,printascii"`
}
