package resources

type VPC struct {
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}

func (v VPC) ResourceType() string {
	return "aws_vpc"
}
