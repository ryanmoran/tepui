package resources

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type AzurermSubnet struct {
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
	VirtualNetworkName string `json:"virtual_network_name"`
	AddressPrefix      string `json:"address_prefix"`
}

func NewAzurermSubnet(name, addressPrefix string, virtualNetwork, resourceGroup terraform.NamedResource) AzurermSubnet {
	return AzurermSubnet{
		Name:               name,
		ResourceGroupName:  resourceGroup.Attribute("name"),
		VirtualNetworkName: virtualNetwork.Attribute("name"),
		AddressPrefix:      addressPrefix,
	}
}
