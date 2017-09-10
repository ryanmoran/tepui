package azure

import "github.com/pivotal-cf/tepui/generate/internal/terraform"

type Template struct {
	Provider  Provider `json:"provider"`
	Resources struct {
		ResourceGroups  terraform.Resources `json:"azurerm_resource_group"`
		VirtualNetworks terraform.Resources `json:"azurerm_virtual_network"`
	} `json:"resource"`
}

func NewTemplate(provider Provider) Template {
	return Template{Provider: provider}
}
