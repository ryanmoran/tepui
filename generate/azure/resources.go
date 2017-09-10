package azure

type ResourceGroup struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (rg ResourceGroup) ResourceType() string {
	return "azurerm_resource_group"
}

type VirtualNetwork struct {
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
	AddressSpace      []string `json:"address_space"`
	Location          string   `json:"location"`
}

func (vn VirtualNetwork) ResourceType() string {
	return "azurerm_virtual_network"
}
