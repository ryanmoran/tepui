package generate

type resource interface {
	resourceType() string
}

type TemplateResourceGoogleComputeNetwork struct {
	Name string `json:"name"`
}

func (trgcn TemplateResourceGoogleComputeNetwork) resourceType() string {
	return "google_compute_network"
}

type TemplateResourceAWSVPC struct {
	CIDRBlock string            `json:"cidr_block"`
	Tags      map[string]string `json:"tags"`
}

func (trav TemplateResourceAWSVPC) resourceType() string {
	return "aws_vpc"
}
