package aws

import (
	"encoding/json"

	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"
)

type TemplateGenerator struct {
	networks NetworkResourceGenerator
}

func NewTemplateGenerator(networks NetworkResourceGenerator) TemplateGenerator {
	return TemplateGenerator{
		networks: networks,
	}
}

func (g TemplateGenerator) Generate(p provider.Provider, m manifest.Manifest) (string, error) {
	template := NewTemplate(Provider{
		AccessKey: p.AWS.AccessKey,
		SecretKey: p.AWS.SecretKey,
		Region:    p.AWS.Region,
	})

	for _, network := range m.Networks {
		template.Resources = append(template.Resources, g.networks.Generate(m.Name, p.AWS.Zones, network)...)
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
