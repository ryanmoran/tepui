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
	template := NewTemplate()
	switch manifest.Provider.Type {
	case "gcp":
		template.Providers.Add(TemplateProviderGoogle{
			Credentials: manifest.Provider.GCP.Credentials,
			Project:     manifest.Provider.GCP.Project,
			Region:      manifest.Provider.GCP.Region,
		})

		for _, network := range manifest.Environment.Networks {
			template.Resources.Add("network", TemplateResourceGoogleComputeNetwork{
				Name: network.Name,
			})
		}

	case "aws":
		template.Providers.Add(TemplateProviderAWS{
			AccessKey: manifest.Provider.AWS.AccessKey,
			SecretKey: manifest.Provider.AWS.SecretKey,
			Region:    manifest.Provider.AWS.Region,
		})

		for _, network := range manifest.Environment.Networks {
			template.Resources.Add("network", TemplateResourceAWSVPC{
				CIDRBlock: network.CIDR,
				Tags: map[string]string{
					"name": network.Name,
				},
			})
		}

	case "azure":
		template.Providers.Add(TemplateProviderAzure{
			SubscriptionID: manifest.Provider.Azure.SubscriptionID,
			ClientID:       manifest.Provider.Azure.ClientID,
			ClientSecret:   manifest.Provider.Azure.ClientSecret,
			TenantID:       manifest.Provider.Azure.TenantID,
		})

		template.Resources.Add("resource_group", TemplateResourceAzureResourceGroup{
			Name:     manifest.Environment.Name,
			Location: manifest.Provider.Azure.Region,
		})

		for _, network := range manifest.Environment.Networks {
			template.Resources.Add("network", TemplateResourceAzureVirtualNetwork{
				Name:              network.Name,
				ResourceGroupName: manifest.Environment.Name,
				AddressSpace:      []string{network.CIDR},
				Location:          manifest.Provider.Azure.Region,
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
