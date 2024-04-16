package groups

type CreateRequest struct {
	Name string `json:"name" db:"name"`
}
