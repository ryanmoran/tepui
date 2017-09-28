package azure

import (
	"github.com/ryanmoran/tepui/generate/azure/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
)

type NetworkResourceGenerator struct{}

func NewNetworkResourceGenerator() NetworkResourceGenerator {
	return NetworkResourceGenerator{}
}

func (g NetworkResourceGenerator) Generate(resourceGroup terraform.NamedResource, region string, network manifest.Network) terraform.Resources {
	var r terraform.Resources

	networkResource := terraform.NamedResource{
		Name:     network.Name,
		Resource: resources.NewAzurermVirtualNetwork(network.Name, network.CIDR, region, resourceGroup),
	}

	r = append(r, networkResource)

	for _, subnet := range network.Subnets {
		subnetResource := terraform.NamedResource{
			Name:     subnet.Name,
			Resource: resources.NewAzurermSubnet(subnet.Name, subnet.CIDR, networkResource, resourceGroup),
		}

		r = append(r, subnetResource)
	}

	return r
}
