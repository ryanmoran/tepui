package azure

import (
	"encoding/json"

	"github.com/ryanmoran/tepui/generate/azure/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
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
		SubscriptionID: p.Azure.SubscriptionID,
		ClientID:       p.Azure.ClientID,
		ClientSecret:   p.Azure.ClientSecret,
		TenantID:       p.Azure.TenantID,
	})

	resourceGroup := terraform.NamedResource{
		Name:     "resource_group",
		Resource: resources.NewAzurermResourceGroup(m.Name, p.Azure.Region),
	}

	template.Resources = append(template.Resources, resourceGroup)

	for _, network := range m.Networks {
		networkResources := g.networks.Generate(resourceGroup, p.Azure.Region, network)
		template.Resources = append(template.Resources, networkResources...)
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
