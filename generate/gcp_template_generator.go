package generate

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/parse"
)

type GCPTemplateGenerator struct{}

func NewGCPTemplateGenerator() GCPTemplateGenerator {
	return GCPTemplateGenerator{}
}

func (g GCPTemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate()
	template.Providers.Add(TemplateProviderGoogle{
		Credentials: provider.GCP.Credentials,
		Project:     provider.GCP.Project,
		Region:      provider.GCP.Region,
	})

	for _, network := range manifest.Networks {
		template.Resources.Add(network.Name, TemplateResourceGoogleComputeNetwork{
			Name: network.Name,
		})
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
