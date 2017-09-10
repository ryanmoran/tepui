package aws

import "encoding/json"

type Resources []NamedResource

func (r Resources) MarshalJSON() ([]byte, error) {
	m := map[string]Resource{}

	for _, nr := range r {
		m[nr.Name] = nr.Resource
	}

	return json.Marshal(m)
}

type NamedResource struct {
	Name     string
	Resource Resource
}

type Resource interface{}

type VPC struct {
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}

func (v VPC) ResourceType() string {
	return "aws_vpc"
}
