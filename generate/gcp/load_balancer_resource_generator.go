package gcp

import (
	"fmt"
	"strings"

	"github.com/ryanmoran/tepui/generate/gcp/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"
)

type LoadBalancerResourceGenerator struct{}

func NewLoadBalancerResourceGenerator() LoadBalancerResourceGenerator {
	return LoadBalancerResourceGenerator{}
}

func (g LoadBalancerResourceGenerator) Generate(loadBalancer manifest.LoadBalancer, zones provider.Zones) terraform.Resources {
	var r terraform.Resources

	globalAddress := terraform.NamedResource{
		Name:     loadBalancer.Name,
		Resource: resources.NewGoogleComputeGlobalAddress(loadBalancer.Name),
	}

	healthCheck := terraform.NamedResource{
		Name:     loadBalancer.Name,
		Resource: resources.NewGoogleComputeHealthCheck(loadBalancer.Name),
	}

	var instanceGroups []terraform.NamedResource
	for _, z := range loadBalancer.Zones {
		zone, _ := zones.Find(z)
		name := strings.Join([]string{loadBalancer.Name, z}, "-")

		instanceGroup := terraform.NamedResource{
			Name:     name,
			Resource: resources.NewGoogleComputeInstanceGroup(name, zone),
		}

		instanceGroups = append(instanceGroups, instanceGroup)
		r = append(r, instanceGroup)
	}

	backendService := terraform.NamedResource{
		Name:     loadBalancer.Name,
		Resource: resources.NewGoogleComputeBackendService(loadBalancer.Name, healthCheck, instanceGroups),
	}

	urlMap := terraform.NamedResource{
		Name:     loadBalancer.Name,
		Resource: resources.NewGoogleComputeUrlMap(loadBalancer.Name, backendService),
	}

	targetHTTPProxy := terraform.NamedResource{
		Name:     loadBalancer.Name,
		Resource: resources.NewGoogleComputeTargetHttpProxy(loadBalancer.Name, urlMap),
	}

	for _, port := range loadBalancer.Ports {
		name := fmt.Sprintf("%s-%d", loadBalancer.Name, port)
		r = append(r, terraform.NamedResource{
			Name:     name,
			Resource: resources.NewGoogleComputeGlobalForwardingRule(name, port, globalAddress, targetHTTPProxy),
		})
	}

	r = append(r, globalAddress, targetHTTPProxy, urlMap, backendService, healthCheck)

	return r
}
