package resources

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type GoogleComputeUrlMap struct {
	Name           string `json:"name"`
	DefaultService string `json:"default_service"`
}

func NewGoogleComputeUrlMap(name string, backendService terraform.NamedResource) GoogleComputeUrlMap {
	return GoogleComputeUrlMap{
		Name:           name,
		DefaultService: backendService.SelfLink(),
	}
}
