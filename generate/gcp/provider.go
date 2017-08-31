package gcp

import "encoding/json"

type Provider struct {
	Credentials string `json:"credentials"`
	Project     string `json:"project"`
	Region      string `json:"region"`
}

func (p Provider) ProviderName() string { return "google" }

func (p Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]map[string]string{
		"google": map[string]string{
			"credentials": p.Credentials,
			"project":     p.Project,
			"region":      p.Region,
		},
	})
}
