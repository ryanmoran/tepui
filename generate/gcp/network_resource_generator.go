package gcp

import (
	"github.com/ryanmoran/tepui/generate/gcp/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
)

type NetworkResourceGenerator struct{}

func NewNetworkResourceGenerator() NetworkResourceGenerator {
	return NetworkResourceGenerator{}
}

func (g NetworkResourceGenerator) Generate(network manifest.Network) terraform.Resources {
	var r terraform.Resources

	networkResource := terraform.NamedResource{
		Name:     network.Name,
		Resource: resources.NewGoogleComputeNetwork(network.Name),
	}

	r = append(r, networkResource)

	for _, subnet := range network.Subnets {
		subnetResource := terraform.NamedResource{
			Name:     subnet.Name,
			Resource: resources.NewGoogleComputeSubnetwork(subnet.Name, subnet.CIDR, networkResource),
		}

		r = append(r, subnetResource)
	}

	return r
}
