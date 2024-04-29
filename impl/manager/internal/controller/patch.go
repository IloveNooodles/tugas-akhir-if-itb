package controller

const (
	PatchReplaceOP = "replace"
	PatchAddOP     = "add"
)

type PatchObject struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
