package azure

import (
	"github.com/pivotal-cf/tepui/generate/azure/resources"
	"github.com/pivotal-cf/tepui/generate/internal/terraform"
	"github.com/pivotal-cf/tepui/parse/manifest"
)

type NetworkResourceGenerator struct{}

func NewNetworkResourceGenerator() NetworkResourceGenerator {
	return NetworkResourceGenerator{}
}

func (g NetworkResourceGenerator) Generate(resourceGroup terraform.NamedResource, region string, network manifest.Network) terraform.Resources {
	var r terraform.Resources

	networkResource := terraform.NamedResource{
		Name: network.Name,
		Resource: resources.AzurermVirtualNetwork{
			Name:              network.Name,
			ResourceGroupName: resourceGroup.Attribute("name"),
			AddressSpace:      []string{network.CIDR},
			Location:          region,
		},
	}

	r = append(r, networkResource)

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

		r = append(r, subnetResource)
	}

	return r
}
