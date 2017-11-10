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

var _ = Describe("LoadBalancerResourceGenerator", func() {
	Describe("Generate", func() {
		It("returns a collection of terraform resources describing a load balancer", func() {
			generator := aws.NewLoadBalancerResourceGenerator()

			loadBalancer := manifest.LoadBalancer{
				Name:  "some-lb",
				Ports: []int{1234, 5678},
				Zones: []string{"az-1", "az-2"},
			}

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

			lbResources := generator.Generate(loadBalancer, zones, "some-vpc-id")
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb",
				Resource: resources.AwsLb{
					Name:             "some-lb",
					LoadBalancerType: "network",
					IdleTimeout:      "3600",
					SecurityGroups:   []string{"${aws_security_group.some-lb.id}"},
				},
			}))
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb-1234",
				Resource: resources.AwsLbTargetGroup{
					Name:     "some-lb-1234",
					Port:     1234,
					Protocol: "http",
					VPCID:    "some-vpc-id",
				},
			}))
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb-5678",
				Resource: resources.AwsLbTargetGroup{
					Name:     "some-lb-5678",
					Port:     5678,
					Protocol: "http",
					VPCID:    "some-vpc-id",
				},
			}))
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb-1234",
				Resource: resources.AwsLbListener{
					LoadBalancerARN: "${aws_lb.some-lb.arn}",
					Port:            "1234",
					DefaultAction: struct {
						TargetGroupARN string `json:"target_group_arn"`
						Type           string `json:"type"`
					}{
						TargetGroupARN: "${aws_lb_target_group.some-lb-1234.arn}",
						Type:           "forward",
					},
				},
			}))
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb-5678",
				Resource: resources.AwsLbListener{
					LoadBalancerARN: "${aws_lb.some-lb.arn}",
					Port:            "5678",
					DefaultAction: struct {
						TargetGroupARN string `json:"target_group_arn"`
						Type           string `json:"type"`
					}{
						TargetGroupARN: "${aws_lb_target_group.some-lb-5678.arn}",
						Type:           "forward",
					},
				},
			}))
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb",
				Resource: resources.AwsSecurityGroup{
					Name: "some-lb",
					Tags: map[string]string{"Name": "some-lb"},
				},
			}))
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb-1234",
				Resource: resources.AwsSecurityGroupRule{
					Type:            "ingress",
					Protocol:        "tcp",
					FromPort:        "1234",
					ToPort:          "1234",
					SecurityGroupID: "${aws_security_group.some-lb.id}",
				},
			}))
			Expect(lbResources).To(ContainElement(terraform.NamedResource{
				Name: "some-lb-5678",
				Resource: resources.AwsSecurityGroupRule{
					Type:            "ingress",
					Protocol:        "tcp",
					FromPort:        "5678",
					ToPort:          "5678",
					SecurityGroupID: "${aws_security_group.some-lb.id}",
				},
			}))
			Expect(lbResources).To(HaveLen(8))
		})
	})
})
