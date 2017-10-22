package resources

import (
	"fmt"

	"github.com/ryanmoran/tepui/generate/internal/terraform"
	"github.com/ryanmoran/tepui/parse/provider"
)

type AwsSubnet struct {
	VPCID            string            `json:"vpc_id"`
	CIDRBlock        string            `json:"cidr_block"`
	AvailabilityZone string            `json:"availability_zone"`
	Tags             map[string]string `json:"tags"`
}

func NewAwsSubnet(name, cidr, environment string, az provider.Zone, vpc terraform.NamedResource) AwsSubnet {
	return AwsSubnet{
		VPCID:            vpc.Attribute("id"),
		CIDRBlock:        cidr,
		AvailabilityZone: az.Name,
		Tags: map[string]string{
			"Name":        fmt.Sprintf("%s-%s", name, az.Alias),
			"Environment": environment,
		},
	}
}
