package gcp

type Provider struct {
	Credentials string `json:"credentials"`
	Project     string `json:"project"`
	Region      string `json:"region"`
}

func (p Provider) ProviderName() string { return "google" }
