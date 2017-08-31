package azure

import "encoding/json"

type ResourceGroupCollection []ResourceGroup

func (rgc ResourceGroupCollection) MarshalJSON() ([]byte, error) {
	m := map[string]ResourceGroup{}

	for _, resourceGroup := range rgc {
		m[resourceGroup.name] = resourceGroup
	}

	return json.Marshal(m)
}

type ResourceGroup struct {
	name string

	Name     string `json:"name"`
	Location string `json:"location"`
}

func (rg ResourceGroup) ResourceType() string {
	return "azurerm_resource_group"
}

type VirtualNetworkCollection []VirtualNetwork

func (vnc VirtualNetworkCollection) MarshalJSON() ([]byte, error) {
	m := map[string]VirtualNetwork{}

	for _, virtualNetwork := range vnc {
		m[virtualNetwork.name] = virtualNetwork
	}

	return json.Marshal(m)
}

type VirtualNetwork struct {
	name string

	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
	AddressSpace      []string `json:"address_space"`
	Location          string   `json:"location"`
}

func (vn VirtualNetwork) ResourceType() string {
	return "azurerm_virtual_network"
}
