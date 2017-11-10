package aws

import (
	"fmt"

	"github.com/ryanmoran/tepui/generate/aws/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"
)

type NetworkResourceGenerator struct{}

func NewNetworkResourceGenerator() NetworkResourceGenerator {
	return NetworkResourceGenerator{}
}

func (g NetworkResourceGenerator) Generate(environment string, availabilityZones []provider.Zone, network manifest.Network) terraform.Resources {
	var r terraform.Resources

	networkResource := terraform.NamedResource{
		Name:     network.Name,
		Resource: resources.NewAwsVpc(network.Name, network.CIDR, environment),
	}

	r = append(r, networkResource)

	for _, subnet := range network.Subnets {
		cidrPartitioner := NewCIDRPartitioner(subnet.CIDR, len(availabilityZones))
		for azIndex, az := range availabilityZones {
			subnetResource := terraform.NamedResource{
				Name:     fmt.Sprintf("%s-%s", subnet.Name, az.Alias),
				Resource: resources.NewAwsSubnet(subnet.Name, cidrPartitioner.Partition(azIndex), environment, az, networkResource),
			}

			r = append(r, subnetResource)
		}
	}

	loadBalancerResourceGenerator := NewLoadBalancerResourceGenerator()
	for _, loadBalancer := range network.LoadBalancers {
		r = append(r, loadBalancerResourceGenerator.Generate(loadBalancer, availabilityZones, networkResource.Attribute("id"))...)
	}

	return r
}
