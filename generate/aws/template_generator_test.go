package aws_test

import (
	"github.com/ryanmoran/tepui/generate/aws"
	"github.com/ryanmoran/tepui/parse/manifest"
	"github.com/ryanmoran/tepui/parse/provider"

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
					Zones: []provider.Zone{
						{
							Alias: "az-1",
							Name:  "us-east-1a",
						},
					},
				},
			}
			manifest := manifest.Manifest{
				Name: "some-environment",
				Networks: []manifest.Network{
					{
						Name: "some-network",
						CIDR: "10.0.0.0/8",
						Subnets: []manifest.Subnet{
							{
								Name: "some-subnet-1",
								CIDR: "10.0.0.0/9",
							},
							{
								Name: "some-subnet-2",
								CIDR: "10.128.0.0/9",
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

			generator := aws.NewTemplateGenerator(aws.NewNetworkResourceGenerator())
			template, err := generator.Generate(prov, manifest)
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
					"aws_lb": {
						"some-lb": {
							"name": "some-lb",
							"load_balancer_type": "network",
							"idle_timeout": "3600",
							"security_groups": ["${aws_security_group.some-lb.id}"]
						}
					},
					"aws_lb_target_group": {
						"some-lb-1234": {
							"name": "some-lb-1234",
							"port": 1234,
							"protocol": "http",
							"vpc_id": "${aws_vpc.some-network.id}"
						}
					},
					"aws_lb_listener": {
						"some-lb-1234": {
							"load_balancer_arn": "${aws_lb.some-lb.arn}",
							"port": "1234",
							"default_action": {
								"target_group_arn": "${aws_lb_target_group.some-lb-1234.arn}",
								"type": "forward"
							}
						}
					},
					"aws_security_group": {
						"some-lb": {
							"name": "some-lb",
							"tags": {
								"Name": "some-lb"
							}
						}
					},
					"aws_security_group_rule": {
						"some-lb-1234": {
							"type": "ingress",
							"protocol": "tcp",
							"from_port": "1234",
							"to_port": "1234",
							"security_group_id": "${aws_security_group.some-lb.id}"
						}
					},
					"aws_vpc": {
						"some-network": {
							"cidr_block": "10.0.0.0/8",
							"tags": {
								"Name": "some-network",
								"Environment": "some-environment"
							}
						}
					},
					"aws_subnet": {
						"some-subnet-1-az-1": {
							"vpc_id": "${aws_vpc.some-network.id}",
							"cidr_block": "10.0.0.0/9",
							"availability_zone": "us-east-1a",
							"tags": {
								"Name": "some-subnet-1-az-1",
								"Environment": "some-environment"
							}
						},
						"some-subnet-2-az-1": {
							"vpc_id": "${aws_vpc.some-network.id}",
							"cidr_block": "10.128.0.0/9",
							"availability_zone": "us-east-1a",
							"tags": {
								"Name": "some-subnet-2-az-1",
								"Environment": "some-environment"
							}
						}
					}
				}
			}`))
		})
	})
})
