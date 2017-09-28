package resources

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type GoogleComputeTargetHttpProxy struct {
	Name   string `json:"name"`
	URLMap string `json:"url_map"`
}

func NewGoogleComputeTargetHttpProxy(name string, urlMap terraform.NamedResource) GoogleComputeTargetHttpProxy {
	return GoogleComputeTargetHttpProxy{
		Name:   name,
		URLMap: urlMap.SelfLink(),
	}
}
