package gcp

import "github.com/pivotal-cf/tepui/generate/internal/terraform"

type Template struct {
	Provider  Provider `json:"provider"`
	Resources struct {
		ComputeNetworks terraform.Resources `json:"google_compute_network"`
	} `json:"resource"`
}

func NewTemplate(provider Provider) Template {
	return Template{Provider: provider}
}
