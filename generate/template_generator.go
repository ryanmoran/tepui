package generate

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (tg TemplateGenerator) Generate(manifest parse.Manifest) (string, error) {
	var template Template
	template.Provider.Google.Credentials = manifest.Provider.GCP.Credentials
	template.Provider.Google.Project = manifest.Provider.GCP.Project
	template.Provider.Google.Region = manifest.Provider.GCP.Region

	template.Resource.GoogleComputeNetwork.Network.Name = manifest.Network.Name

	output, err := json.Marshal(template)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

type Template struct {
	Provider TemplateProvider `json:"provider"`
	Resource TemplateResource `json:"resource"`
}

type TemplateProvider struct {
	Google TemplateProviderGoogle `json:"google"`
}

type TemplateProviderGoogle struct {
	Credentials string `json:"credentials"`
	Project     string `json:"project"`
	Region      string `json:"region"`
}

type TemplateResource struct {
	GoogleComputeNetwork TemplateResourceGoogleComputeNetwork `json:"google_compute_network"`
}

type TemplateResourceGoogleComputeNetwork struct {
	Network struct {
		Name string `json:"name"`
	} `json:"network"`
}
