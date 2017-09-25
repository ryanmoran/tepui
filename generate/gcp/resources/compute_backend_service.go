package resources

type GoogleComputeBackendService struct {
	Name         string                               `json:"name"`
	Backend      []GoogleComputeBackendServiceBackend `json:"backend"`
	HealthChecks []string                             `json:"health_checks"`
}

type GoogleComputeBackendServiceBackend struct {
	Group string `json:"group"`
}
