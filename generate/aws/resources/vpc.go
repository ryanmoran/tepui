package resources

type VPC struct {
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}
