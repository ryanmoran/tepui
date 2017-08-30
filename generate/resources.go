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

type TemplateResourceAzureResourceGroup struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (trarg TemplateResourceAzureResourceGroup) resourceType() string {
	return "azurerm_resource_group"
}

type TemplateResourceAzureVirtualNetwork struct {
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
	AddressSpace      []string `json:"address_space"`
	Location          string   `json:"location"`
}

func (travn TemplateResourceAzureVirtualNetwork) resourceType() string {
	return "azurerm_virtual_network"
}
