package resources

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type GoogleComputeBackendService struct {
	Name         string                               `json:"name"`
	Backend      []GoogleComputeBackendServiceBackend `json:"backend"`
	HealthChecks []string                             `json:"health_checks"`
}

func NewGoogleComputeBackendService(name string, healthCheck terraform.NamedResource, instanceGroups []terraform.NamedResource) GoogleComputeBackendService {
	var backends []GoogleComputeBackendServiceBackend
	for _, instanceGroup := range instanceGroups {
		backend := GoogleComputeBackendServiceBackend{
			Group: instanceGroup.SelfLink(),
		}
		backends = append(backends, backend)
	}

	return GoogleComputeBackendService{
		Name:    name,
		Backend: backends,
		HealthChecks: []string{
			healthCheck.SelfLink(),
		},
	}
}

type GoogleComputeBackendServiceBackend struct {
	Group string `json:"group"`
}
