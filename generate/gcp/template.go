package gcp

type Template struct {
	Provider  Provider `json:"provider"`
	Resources struct {
		ComputeNetworks ComputeNetworksCollection `json:"google_compute_network"`
	} `json:"resource"`
}

func NewTemplate(provider Provider) Template {
	return Template{Provider: provider}
}
