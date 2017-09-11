package aws

import (
	"github.com/pivotal-cf/tepui/generate/aws/resources"
	"github.com/pivotal-cf/tepui/generate/internal/terraform"
	"github.com/pivotal-cf/tepui/parse/manifest"
)

type NetworkResourceGenerator struct{}

func NewNetworkResourceGenerator() NetworkResourceGenerator {
	return NetworkResourceGenerator{}
}

func (g NetworkResourceGenerator) Generate(environment string, network manifest.Network) terraform.Resources {
	var r terraform.Resources

	networkResource := terraform.NamedResource{
		Name: network.Name,
		Resource: resources.AwsVpc{
			CIDRBlock: network.CIDR,
			Tags: map[string]string{
				"Name":        network.Name,
				"Environment": environment,
			},
		},
	}

	r = append(r, networkResource)

	for _, subnet := range network.Subnets {
		subnetResource := terraform.NamedResource{
			Name: subnet.Name,
			Resource: resources.AwsSubnet{
				VPCID:     networkResource.Attribute("id"),
				CIDRBlock: subnet.CIDR,
				Tags: map[string]string{
					"Name":        subnet.Name,
					"Environment": environment,
				},
			},
		}

		r = append(r, subnetResource)
	}

	return r
}
