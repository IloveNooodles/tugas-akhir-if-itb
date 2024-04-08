package dto

type SuccessResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
