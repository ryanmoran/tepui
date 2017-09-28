package aws

import (
	"github.com/ryanmoran/tepui/generate/aws/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
)

type NetworkResourceGenerator struct{}

func NewNetworkResourceGenerator() NetworkResourceGenerator {
	return NetworkResourceGenerator{}
}

func (g NetworkResourceGenerator) Generate(environment string, network manifest.Network) terraform.Resources {
	var r terraform.Resources

	networkResource := terraform.NamedResource{
		Name:     network.Name,
		Resource: resources.NewAwsVpc(network.Name, network.CIDR, environment),
	}

	r = append(r, networkResource)

	for _, subnet := range network.Subnets {
		subnetResource := terraform.NamedResource{
			Name:     subnet.Name,
			Resource: resources.NewAwsSubnet(subnet.Name, subnet.CIDR, environment, networkResource),
		}

		r = append(r, subnetResource)
	}

	return r
}
