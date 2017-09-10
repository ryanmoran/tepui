package resources

type ResourceGroup struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (rg ResourceGroup) ResourceType() string {
	return "azurerm_resource_group"
}
