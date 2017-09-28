package aws_test

import (
	"github.com/ryanmoran/tepui/generate/aws"
	"github.com/ryanmoran/tepui/generate/aws/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NetworkResourceGenerator", func() {
	Describe("Generate", func() {
		It("returns a collection of terraform resources describing a network", func() {
			generator := aws.NewNetworkResourceGenerator()

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

			Expect(generator.Generate("some-environment", network)).To(ConsistOf(terraform.Resources{
				{
					Name: "some-network",
					Resource: resources.AwsVpc{
						CIDRBlock: "1.2.3.4/5",
						Tags: map[string]string{
							"Name":        "some-network",
							"Environment": "some-environment",
						},
					},
				},
				{
					Name: "some-subnet",
					Resource: resources.AwsSubnet{
						VPCID:     "${aws_vpc.some-network.id}",
						CIDRBlock: "6.7.8.9/10",
						Tags: map[string]string{
							"Name":        "some-subnet",
							"Environment": "some-environment",
						},
					},
				},
			}))
		})
	})
})
