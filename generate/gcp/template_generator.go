package gcp

import (
	"encoding/json"

	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"
)

type TemplateGenerator struct {
	networks      NetworkResourceGenerator
	loadBalancers LoadBalancerResourceGenerator
}

func NewTemplateGenerator(networks NetworkResourceGenerator, loadBalancers LoadBalancerResourceGenerator) TemplateGenerator {
	return TemplateGenerator{
		networks:      networks,
		loadBalancers: loadBalancers,
	}
}

func (g TemplateGenerator) Generate(p provider.Provider, m manifest.Manifest) (string, error) {
	template := NewTemplate(Provider{
		Credentials: p.GCP.Credentials,
		Project:     p.GCP.Project,
		Region:      p.GCP.Region,
	})

	for _, network := range m.Networks {
		networkResources := g.networks.Generate(network, p.GCP.Zones)
		template.Resources = append(template.Resources, networkResources...)
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
