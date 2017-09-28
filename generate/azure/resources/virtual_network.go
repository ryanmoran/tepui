package resources

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type AzurermVirtualNetwork struct {
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
	AddressSpace      []string `json:"address_space"`
	Location          string   `json:"location"`
}

func NewAzurermVirtualNetwork(name, addressSpace, location string, resourceGroup terraform.NamedResource) AzurermVirtualNetwork {
	return AzurermVirtualNetwork{
		Name:              name,
		ResourceGroupName: resourceGroup.Attribute("name"),
		AddressSpace:      []string{addressSpace},
		Location:          location,
	}
}
