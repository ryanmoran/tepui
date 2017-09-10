package azure

type Template struct {
	Provider  Provider `json:"provider"`
	Resources struct {
		ResourceGroups  Resources `json:"azurerm_resource_group"`
		VirtualNetworks Resources `json:"azurerm_virtual_network"`
	} `json:"resource"`
}

func NewTemplate(provider Provider) Template {
	return Template{Provider: provider}
}
