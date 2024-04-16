package company

type CreateRequest struct {
	Name        string `json:"name" validate:"required,printascii"`
	ClusterName string `json:"cluster_name" validate:"required,printascii"`
}
