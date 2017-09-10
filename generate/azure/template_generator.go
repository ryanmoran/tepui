package azure

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (g TemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate(Provider{
		SubscriptionID: provider.Azure.SubscriptionID,
		ClientID:       provider.Azure.ClientID,
		ClientSecret:   provider.Azure.ClientSecret,
		TenantID:       provider.Azure.TenantID,
	})

	template.Resources.ResourceGroups = append(template.Resources.ResourceGroups, NamedResource{
		Name: "resource_group",
		Resource: ResourceGroup{
			Name:     manifest.Name,
			Location: provider.Azure.Region,
		},
	})

	for _, network := range manifest.Networks {
		template.Resources.VirtualNetworks = append(template.Resources.VirtualNetworks, NamedResource{
			Name: network.Name,
			Resource: VirtualNetwork{
				Name:              network.Name,
				ResourceGroupName: "${azurerm_resource_group.resource_group.name}",
				AddressSpace:      []string{network.CIDR},
				Location:          provider.Azure.Region,
			},
		})
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
