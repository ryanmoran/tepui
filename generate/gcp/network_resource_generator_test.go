package gcp_test

import (
	"github.com/ryanmoran/tepui/generate/gcp"
	"github.com/ryanmoran/tepui/generate/gcp/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NetworkResourceGenerator", func() {
	Describe("Generate", func() {
		It("returns a collection of terraform resources describing a network", func() {
			generator := gcp.NewNetworkResourceGenerator()

			zones := []provider.Zone{
				{
					Alias: "az-1",
					Name:  "some-zone-1",
				},
				{
					Alias: "az-2",
					Name:  "some-zone-2",
				},
			}

			network := manifest.Network{
				Name: "some-network",
				CIDR: "1.2.3.4/5",
				Subnets: []manifest.Subnet{
					{
						Name: "some-subnet",
						CIDR: "6.7.8.9/10",
					},
				},
				LoadBalancers: []manifest.LoadBalancer{
					{
						Name:  "some-lb",
						Ports: []int{1234, 5678},
						Zones: []string{"az-1", "az-2"},
					},
				},
			}

			Expect(generator.Generate(network, zones)).To(ConsistOf(terraform.Resources{
				{
					Name: "some-network",
					Resource: resources.GoogleComputeNetwork{
						Name: "some-network",
					},
				},
				{
					Name: "some-subnet",
					Resource: resources.GoogleComputeSubnetwork{
						Name:        "some-subnet",
						IPCIDRRange: "6.7.8.9/10",
						Network:     "${google_compute_network.some-network.self_link}",
					},
				},
				{
					Name: "some-lb",
					Resource: resources.GoogleComputeGlobalAddress{
						Name: "some-lb",
					},
				},
				{
					Name: "some-lb-1234",
					Resource: resources.GoogleComputeGlobalForwardingRule{
						Name:      "some-lb-1234",
						IPAddress: "${google_compute_global_address.some-lb.address}",
						PortRange: "1234",
						Target:    "${google_compute_target_http_proxy.some-lb.self_link}",
					},
				},
				{
					Name: "some-lb-5678",
					Resource: resources.GoogleComputeGlobalForwardingRule{
						Name:      "some-lb-5678",
						IPAddress: "${google_compute_global_address.some-lb.address}",
						PortRange: "5678",
						Target:    "${google_compute_target_http_proxy.some-lb.self_link}",
					},
				},
				{
					Name: "some-lb",
					Resource: resources.GoogleComputeTargetHttpProxy{
						Name:   "some-lb",
						URLMap: "${google_compute_url_map.some-lb.self_link}",
					},
				},
				{
					Name: "some-lb",
					Resource: resources.GoogleComputeUrlMap{
						Name:           "some-lb",
						DefaultService: "${google_compute_backend_service.some-lb.self_link}",
					},
				},
				{
					Name: "some-lb",
					Resource: resources.GoogleComputeBackendService{
						Name: "some-lb",
						Backend: []resources.GoogleComputeBackendServiceBackend{
							{
								Group: "${google_compute_instance_group.some-lb-az-1.self_link}",
							},
							{
								Group: "${google_compute_instance_group.some-lb-az-2.self_link}",
							},
						},
						HealthChecks: []string{
							"${google_compute_health_check.some-lb.self_link}",
						},
					},
				},
				{
					Name: "some-lb-az-1",
					Resource: resources.GoogleComputeInstanceGroup{
						Name:    "some-lb-az-1",
						Zone:    "some-zone-1",
						Network: "${google_compute_network.some-network.self_link}",
					},
				},
				{
					Name: "some-lb-az-2",
					Resource: resources.GoogleComputeInstanceGroup{
						Name:    "some-lb-az-2",
						Zone:    "some-zone-2",
						Network: "${google_compute_network.some-network.self_link}",
					},
				},
				{
					Name: "some-lb",
					Resource: resources.GoogleComputeHealthCheck{
						Name:           "some-lb",
						TCPHealthCheck: resources.GoogleComputeHealthCheckTCP{},
					},
				},
			}))
		})
	})
})
