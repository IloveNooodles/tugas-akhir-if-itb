package deployment

type Deployment struct {
	Version  string   `json:"version" validate:"required,printascii"`
	Name     string   `json:"name" validate:"required,printascii"`
	Category string   `json:"category" validate:"required,printascii"`
	Steps    []string `json:"steps"`
	RecipeID string   `json:"recipeID" validate:"omitempty,printascii"`
}

type Recipes struct {
}

type Action struct {
	Name string `json:"name" validate:"required,printascii"`
	Type string `json:"type" validate:"required,printascii"`
}

type Input struct{}

type Output struct{}
