package gcp

import (
	"fmt"
	"strconv"
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
		Name: loadBalancer.Name,
		Resource: resources.GoogleComputeGlobalAddress{
			Name: loadBalancer.Name,
		},
	}

	zone, _ := zones.Find(loadBalancer.Zones[0])

	instanceGroup := terraform.NamedResource{
		Name: strings.Join([]string{loadBalancer.Name, loadBalancer.Zones[0]}, "-"),
		Resource: resources.GoogleComputeInstanceGroup{
			Name: strings.Join([]string{loadBalancer.Name, loadBalancer.Zones[0]}, "-"),
			Zone: zone,
		},
	}

	healthCheck := terraform.NamedResource{
		Name: loadBalancer.Name,
		Resource: resources.GoogleComputeHealthCheck{
			Name:           loadBalancer.Name,
			TCPHealthCheck: resources.GoogleComputeHealthCheckTCP{},
		},
	}

	backendService := terraform.NamedResource{
		Name: loadBalancer.Name,
		Resource: resources.GoogleComputeBackendService{
			Name: loadBalancer.Name,
			Backend: []resources.GoogleComputeBackendServiceBackend{
				{
					Group: instanceGroup.SelfLink(),
				},
			},
			HealthChecks: []string{
				healthCheck.SelfLink(),
			},
		},
	}

	urlMap := terraform.NamedResource{
		Name: loadBalancer.Name,
		Resource: resources.GoogleComputeUrlMap{
			Name:           loadBalancer.Name,
			DefaultService: backendService.SelfLink(),
		},
	}

	targetHTTPProxy := terraform.NamedResource{
		Name: loadBalancer.Name,
		Resource: resources.GoogleComputeTargetHttpProxy{
			Name:   loadBalancer.Name,
			URLMap: urlMap.SelfLink(),
		},
	}

	for _, port := range loadBalancer.Ports {
		forwardingRule := terraform.NamedResource{
			Name: fmt.Sprintf("%s-%d", loadBalancer.Name, port),
			Resource: resources.GoogleComputeGlobalForwardingRule{
				Name:      fmt.Sprintf("%s-%d", loadBalancer.Name, port),
				IPAddress: globalAddress.Attribute("address"),
				PortRange: strconv.Itoa(port),
				Target:    targetHTTPProxy.SelfLink(),
			},
		}

		r = append(r, forwardingRule)
	}

	r = append(r, globalAddress, targetHTTPProxy, urlMap, backendService, healthCheck, instanceGroup)

	return r
}
