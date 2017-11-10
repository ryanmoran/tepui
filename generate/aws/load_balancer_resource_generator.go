package aws

import (
	"fmt"

	"github.com/ryanmoran/tepui/generate/aws/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"
)

type LoadBalancerResourceGenerator struct{}

func NewLoadBalancerResourceGenerator() LoadBalancerResourceGenerator {
	return LoadBalancerResourceGenerator{}
}

func (g LoadBalancerResourceGenerator) Generate(loadBalancer manifest.LoadBalancer, zones provider.Zones, vpcID string) terraform.Resources {
	var r terraform.Resources

	securityGroup := terraform.NamedResource{
		Name:     loadBalancer.Name,
		Resource: resources.NewAwsSecurityGroup(loadBalancer.Name),
	}
	r = append(r, securityGroup)

	for _, port := range loadBalancer.Ports {
		r = append(r, terraform.NamedResource{
			Name:     fmt.Sprintf("%s-%d", loadBalancer.Name, port),
			Resource: resources.NewAwsSecurityGroupRule(securityGroup.Attribute("id"), port),
		})
	}

	lb := terraform.NamedResource{
		Name:     loadBalancer.Name,
		Resource: resources.NewAwsLb(loadBalancer.Name, securityGroup.Attribute("id")),
	}
	r = append(r, lb)

	for _, port := range loadBalancer.Ports {
		name := fmt.Sprintf("%s-%d", loadBalancer.Name, port)

		tg := terraform.NamedResource{
			Name:     name,
			Resource: resources.NewAwsLbTargetGroup(name, vpcID, port),
		}

		l := terraform.NamedResource{
			Name:     name,
			Resource: resources.NewAwsLbListener(lb.Attribute("arn"), tg.Attribute("arn"), port),
		}

		r = append(r, l, tg)
	}

	return r
}
