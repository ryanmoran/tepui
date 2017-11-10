package gcp

import (
	"github.com/ryanmoran/tepui/generate/gcp/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"
)

type NetworkResourceGenerator struct{}

func NewNetworkResourceGenerator() NetworkResourceGenerator {
	return NetworkResourceGenerator{}
}

func (g NetworkResourceGenerator) Generate(network manifest.Network, zones provider.Zones) terraform.Resources {
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

	loadBalancerResourceGenerator := NewLoadBalancerResourceGenerator()
	for _, loadBalancer := range network.LoadBalancers {
		r = append(r, loadBalancerResourceGenerator.Generate(loadBalancer, zones, networkResource.SelfLink())...)
	}

	return r
}
