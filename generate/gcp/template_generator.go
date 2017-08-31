package gcp

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/generate"
	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (g TemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := generate.NewTemplate()
	template.Providers.Add(generate.TemplateProviderGoogle{
		Credentials: provider.GCP.Credentials,
		Project:     provider.GCP.Project,
		Region:      provider.GCP.Region,
	})

	for _, network := range manifest.Networks {
		template.Resources.Add(network.Name, generate.TemplateResourceGoogleComputeNetwork{
			Name: network.Name,
		})
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
