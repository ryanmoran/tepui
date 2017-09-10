package azure

import (
	"encoding/json"

	"github.com/pivotal-cf/tepui/generate/azure/resources"
	"github.com/pivotal-cf/tepui/generate/internal/terraform"
	"github.com/pivotal-cf/tepui/parse"
	"github.com/pivotal-cf/tepui/parse/provider"
)

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (g TemplateGenerator) Generate(prov provider.Provider, manifest parse.Manifest) (string, error) {
	template := NewTemplate(Provider{
		SubscriptionID: prov.Azure.SubscriptionID,
		ClientID:       prov.Azure.ClientID,
		ClientSecret:   prov.Azure.ClientSecret,
		TenantID:       prov.Azure.TenantID,
	})

	resourceGroup := terraform.NamedResource{
		Name: "resource_group",
		Resource: resources.AzurermResourceGroup{
			Name:     manifest.Name,
			Location: prov.Azure.Region,
		},
	}

	template.Resources = append(template.Resources, resourceGroup)

	for _, network := range manifest.Networks {
		networkResource := terraform.NamedResource{
			Name: network.Name,
			Resource: resources.AzurermVirtualNetwork{
				Name:              network.Name,
				ResourceGroupName: resourceGroup.Attribute("name"),
				AddressSpace:      []string{network.CIDR},
				Location:          prov.Azure.Region,
			},
		}

		template.Resources = append(template.Resources, networkResource)

		for _, subnet := range network.Subnets {
			subnetResource := terraform.NamedResource{
				Name: subnet.Name,
				Resource: resources.AzurermSubnet{
					Name:               subnet.Name,
					ResourceGroupName:  resourceGroup.Attribute("name"),
					VirtualNetworkName: networkResource.Attribute("name"),
					AddressPrefix:      subnet.CIDR,
				},
			}

			template.Resources = append(template.Resources, subnetResource)
		}
	}

	output, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
