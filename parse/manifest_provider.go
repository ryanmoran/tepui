package parse

import "fmt"

type ManifestProvider struct {
	Type string
	GCP  ManifestProviderGCP
}

func (mpp *ManifestProvider) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var provider struct {
		Type   string `yaml:"type"`
		Params map[string]interface{}
	}

	err := unmarshal(&provider)
	if err != nil {
		return err
	}

	switch provider.Type {
	case "gcp":
		mpp.Type = provider.Type
		mpp.GCP.Credentials = provider.Params["credentials"].(string)
		mpp.GCP.Project = provider.Params["project"].(string)
		mpp.GCP.Region = provider.Params["region"].(string)
	default:
		return fmt.Errorf("unknown provider type: %q", provider.Type)
	}

	return nil
}

type ManifestProviderGCP struct {
	Credentials string
	Project     string
	Region      string
}
