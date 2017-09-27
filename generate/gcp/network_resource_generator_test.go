package gcp_test

import (
	"github.com/ryanmoran/tepui/generate/gcp"
	"github.com/ryanmoran/tepui/generate/gcp/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NetworkResourceGenerator", func() {
	Describe("Generate", func() {
		It("returns a collection of terraform resources describing a network", func() {
			generator := gcp.NewNetworkResourceGenerator()

			network := manifest.Network{
				Name: "some-network",
				CIDR: "1.2.3.4/5",
				Subnets: []manifest.Subnet{
					{
						Name: "some-subnet",
						CIDR: "6.7.8.9/10",
					},
				},
			}

			Expect(generator.Generate(network)).To(ConsistOf(terraform.Resources{
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
			}))
		})
	})
})
