package generate

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/parse"
)

type AzureTemplateGenerator struct{}

func NewAzureTemplateGenerator() AzureTemplateGenerator {
	return AzureTemplateGenerator{}
}

func (g AzureTemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate()
	template.Providers.Add(TemplateProviderAzure{
		SubscriptionID: provider.Azure.SubscriptionID,
		ClientID:       provider.Azure.ClientID,
		ClientSecret:   provider.Azure.ClientSecret,
		TenantID:       provider.Azure.TenantID,
	})

	template.Resources.Add("resource_group", TemplateResourceAzureResourceGroup{
		Name:     manifest.Name,
		Location: provider.Azure.Region,
	})

	for _, network := range manifest.Networks {
		template.Resources.Add(network.Name, TemplateResourceAzureVirtualNetwork{
			Name:              network.Name,
			ResourceGroupName: "${azurerm_resource_group.resource_group.name}",
			AddressSpace:      []string{network.CIDR},
			Location:          provider.Azure.Region,
		})
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
