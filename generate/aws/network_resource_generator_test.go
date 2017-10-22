package aws_test

import (
	"github.com/ryanmoran/tepui/generate/aws"
	"github.com/ryanmoran/tepui/generate/aws/resources"
	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NetworkResourceGenerator", func() {
	Describe("Generate", func() {
		It("returns a collection of terraform resources describing a network", func() {
			generator := aws.NewNetworkResourceGenerator()

			network := manifest.Network{
				Name: "some-network",
				CIDR: "127.0.0.0/8",
				Subnets: []manifest.Subnet{
					{
						Name: "some-subnet",
						CIDR: "127.0.0.0/8",
					},
				},
			}

			zones := []provider.Zone{
				{
					Name:  "us-central-1a",
					Alias: "some-az-1",
				},
				{
					Name:  "us-central-1b",
					Alias: "some-az-2",
				},
			}

			Expect(generator.Generate("some-environment", zones, network)).To(ConsistOf(terraform.Resources{
				{
					Name: "some-network",
					Resource: resources.AwsVpc{
						CIDRBlock: "127.0.0.0/8",
						Tags: map[string]string{
							"Name":        "some-network",
							"Environment": "some-environment",
						},
					},
				},
				{
					Name: "some-subnet-some-az-1",
					Resource: resources.AwsSubnet{
						VPCID:            "${aws_vpc.some-network.id}",
						CIDRBlock:        "127.0.0.0/9",
						AvailabilityZone: "us-central-1a",
						Tags: map[string]string{
							"Name":        "some-subnet-some-az-1",
							"Environment": "some-environment",
						},
					},
				},
				{
					Name: "some-subnet-some-az-2",
					Resource: resources.AwsSubnet{
						VPCID:            "${aws_vpc.some-network.id}",
						CIDRBlock:        "127.128.0.0/9",
						AvailabilityZone: "us-central-1b",
						Tags: map[string]string{
							"Name":        "some-subnet-some-az-2",
							"Environment": "some-environment",
						},
					},
				},
			}))
		})
	})
})
