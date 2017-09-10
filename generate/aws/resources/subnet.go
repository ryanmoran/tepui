package resources

type AwsSubnet struct {
	VPCID     string            `json:"vpc_id"`
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}
