package gcp_test

import (
	"github.com/ryanmoran/tepui/generate/gcp"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateGenerator", func() {
	Describe("Generate", func() {
		It("generates a template from the given manifest", func() {
			prov := provider.Provider{
				Type: "gcp",
				GCP: provider.GCP{
					Credentials: "some-credentials",
					Project:     "some-project",
					Region:      "some-region",
					Zones: []provider.Zone{
						{
							Alias: "az-1",
							Name:  "some-zone",
						},
					},
				},
			}
			manifest := manifest.Manifest{
				Name: "some-environment",
				Networks: []manifest.Network{
					{
						Name: "some-network",
						CIDR: "1.2.3.4/5",
						Subnets: []manifest.Subnet{
							{
								Name: "some-subnet",
								CIDR: "2.3.4.5/6",
							},
						},
						LoadBalancers: []manifest.LoadBalancer{
							{
								Name:  "some-lb",
								Ports: []int{1234},
								Zones: []string{"az-1"},
							},
						},
					},
				},
			}

			generator := gcp.NewTemplateGenerator(gcp.NewNetworkResourceGenerator(), gcp.NewLoadBalancerResourceGenerator())
			template, err := generator.Generate(prov, manifest)
			Expect(err).NotTo(HaveOccurred())
			Expect(template).To(MatchJSON(`{
				"provider": {
					"google": {
						"credentials": "some-credentials",
						"project": "some-project",
						"region": "some-region"
					}
				},
				"resource": {
					"google_compute_network": {
						"some-network": {
							"name": "some-network"
						} },
					"google_compute_subnetwork": {
						"some-subnet": {
							"name": "some-subnet",
							"ip_cidr_range": "2.3.4.5/6",
							"network": "${google_compute_network.some-network.self_link}"
						}
					},
					"google_compute_global_address": {
						"some-lb": {
							"name": "some-lb"
						}
					},
					"google_compute_global_forwarding_rule": {
						"some-lb-1234": {
							"ip_address": "${google_compute_global_address.some-lb.address}",
							"name": "some-lb-1234",
							"port_range": "1234",
							"target": "${google_compute_target_http_proxy.some-lb.self_link}"
						}
					},
					"google_compute_target_http_proxy": {
						"some-lb": {
							"name": "some-lb",
							"url_map": "${google_compute_url_map.some-lb.self_link}"
						}
					},
					"google_compute_url_map": {
						"some-lb": {
							"default_service": "${google_compute_backend_service.some-lb.self_link}",
							"name": "some-lb"
						}
					},
					"google_compute_backend_service": {
						"some-lb": {
							"backend": [{
								"group": "${google_compute_instance_group.some-lb-az-1.self_link}"
							}],
							"health_checks": [
								"${google_compute_health_check.some-lb.self_link}"
							],
							"name": "some-lb"
						}
					},
					"google_compute_instance_group": {
						"some-lb-az-1": {
							"name": "some-lb-az-1",
							"zone": "some-zone",
							"network": "${google_compute_network.some-network.self_link}"
						}
					},
					"google_compute_health_check": {
						"some-lb": {
							"name": "some-lb",
							"tcp_health_check": {}
						}
					}
				}
			}`))
		})
	})
})
