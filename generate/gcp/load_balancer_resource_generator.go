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

	healthCheck := terraform.NamedResource{
		Name: loadBalancer.Name,
		Resource: resources.GoogleComputeHealthCheck{
			Name:           loadBalancer.Name,
			TCPHealthCheck: resources.GoogleComputeHealthCheckTCP{},
		},
	}

	var instanceGroups []terraform.NamedResource
	for _, z := range loadBalancer.Zones {
		zone, _ := zones.Find(z)

		instanceGroup := terraform.NamedResource{
			Name: strings.Join([]string{loadBalancer.Name, z}, "-"),
			Resource: resources.GoogleComputeInstanceGroup{
				Name: strings.Join([]string{loadBalancer.Name, z}, "-"),
				Zone: zone,
			},
		}

		instanceGroups = append(instanceGroups, instanceGroup)
		r = append(r, instanceGroup)
	}

	var backends []resources.GoogleComputeBackendServiceBackend
	for _, instanceGroup := range instanceGroups {
		backend := resources.GoogleComputeBackendServiceBackend{
			Group: instanceGroup.SelfLink(),
		}
		backends = append(backends, backend)
	}

	backendService := terraform.NamedResource{
		Name: loadBalancer.Name,
		Resource: resources.GoogleComputeBackendService{
			Name:    loadBalancer.Name,
			Backend: backends,
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

	r = append(r, globalAddress, targetHTTPProxy, urlMap, backendService, healthCheck)

	return r
}
