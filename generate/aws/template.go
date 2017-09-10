package aws

import "github.com/pivotal-cf/tepui/generate/internal/terraform"

type Template struct {
	Provider  Provider `json:"provider"`
	Resources struct {
		VPCs terraform.Resources `json:"aws_vpc"`
	} `json:"resource"`
}

func NewTemplate(provider Provider) Template {
	return Template{Provider: provider}
}
