package aws_test

import (
	"github.com/pivotal-cf/tepui/generate/aws"
	"github.com/pivotal-cf/tepui/parse/manifest"
	"github.com/pivotal-cf/tepui/parse/provider"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateGenerator", func() {
	Describe("Generate", func() {
		It("generates a template from the given manifest", func() {
			prov := provider.Provider{
				Type: "aws",
				AWS: provider.AWS{
					AccessKey: "some-access-key",
					SecretKey: "some-secret-key",
					Region:    "some-region",
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

			template, err := aws.NewTemplateGenerator().Generate(prov, manifest)
			Expect(err).NotTo(HaveOccurred())
			Expect(template).To(MatchJSON(`{
				"provider": {
					"aws": {
						"access_key": "some-access-key",
						"secret_key": "some-secret-key",
						"region": "some-region"
					}
				},
				"resource": {
					"aws_vpc": {
						"some-network": {
							"cidr_block": "1.2.3.4/5",
							"tags": {
								"Name": "some-network",
								"Environment": "some-environment"
							}
						}
					},
					"aws_subnet": {
						"some-subnet": {
							"vpc_id": "${aws_vpc.some-network.id}",
							"cidr_block": "2.3.4.5/6",
							"tags": {
								"Name": "some-subnet",
								"Environment": "some-environment"
							}
						}
					}
				}
			}`))
		})
	})
})
