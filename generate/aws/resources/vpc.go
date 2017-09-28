package resources

type AwsVpc struct {
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}

func NewAwsVpc(name, cidr, environment string) AwsVpc {
	return AwsVpc{
		CIDRBlock: cidr,
		Tags: map[string]string{
			"Name":        name,
			"Environment": environment,
		},
	}
}
