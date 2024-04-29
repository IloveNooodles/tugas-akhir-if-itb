package controller

type DeployParams struct {
	Replica int32
	Name    string
	Image   string
	Labels  map[string]string
	Targets map[string]string
}
