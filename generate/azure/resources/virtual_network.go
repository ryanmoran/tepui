package resources

type VirtualNetwork struct {
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
	AddressSpace      []string `json:"address_space"`
	Location          string   `json:"location"`
}

func (vn VirtualNetwork) ResourceType() string {
	return "azurerm_virtual_network"
}
