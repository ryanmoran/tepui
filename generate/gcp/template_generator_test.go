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
					},
				},
			}

			generator := gcp.NewTemplateGenerator(gcp.NewNetworkResourceGenerator())
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
						}
					},
					"google_compute_subnetwork": {
						"some-subnet": {
							"name": "some-subnet",
							"ip_cidr_range": "2.3.4.5/6",
							"network": "${google_compute_network.some-network.self_link}"
						}
					}
				}
			}`))
		})
	})
})
