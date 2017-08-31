package aws

import "encoding/json"

type VPCCollection []VPC

func (vc VPCCollection) MarshalJSON() ([]byte, error) {
	m := map[string]VPC{}

	for _, vpc := range vc {
		m[vpc.name] = vpc
	}

	return json.Marshal(m)
}

type VPC struct {
	name string

	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}

func (v VPC) ResourceType() string {
	return "aws_vpc"
}
