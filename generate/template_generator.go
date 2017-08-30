package generate

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/parse"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (tg TemplateGenerator) Generate(provider parse.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate()
	switch provider.Type {
	case "gcp":
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

	case "aws":
		template.Providers.Add(TemplateProviderAWS{
			AccessKey: provider.AWS.AccessKey,
			SecretKey: provider.AWS.SecretKey,
			Region:    provider.AWS.Region,
		})

		for _, network := range manifest.Networks {
			template.Resources.Add(network.Name, TemplateResourceAWSVPC{
				CIDRBlock: network.CIDR,
				Tags: map[string]string{
					"name": network.Name,
				},
			})
		}

	case "azure":
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

	default:
		panic("unknown provider")
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
