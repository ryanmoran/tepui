package azure

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type Template struct {
	Provider  Provider            `json:"provider"`
	Resources terraform.Resources `json:"resource"`
}

func NewTemplate(provider Provider) Template {
	return Template{Provider: provider}
}
