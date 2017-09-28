package resources

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type AwsSubnet struct {
	VPCID     string            `json:"vpc_id"`
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}

func NewAwsSubnet(name, cidr, environment string, vpc terraform.NamedResource) AwsSubnet {
	return AwsSubnet{
		VPCID:     vpc.Attribute("id"),
		CIDRBlock: cidr,
		Tags: map[string]string{
			"Name":        name,
			"Environment": environment,
		},
	}
}
