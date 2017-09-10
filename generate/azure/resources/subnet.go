package resources

type AzurermSubnet struct {
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
	VirtualNetworkName string `json:"virtual_network_name"`
	AddressPrefix      string `json:"address_prefix"`
}
