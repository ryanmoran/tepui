package parse

import "fmt"

type ManifestProvider struct {
	Type string
	GCP  ManifestProviderGCP
	AWS  ManifestProviderAWS
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

	mpp.Type = provider.Type

	switch provider.Type {
	case "gcp":
		mpp.GCP.Credentials = provider.Params["credentials"].(string)
		mpp.GCP.Project = provider.Params["project"].(string)
		mpp.GCP.Region = provider.Params["region"].(string)
	case "aws":
		mpp.AWS.AccessKey = provider.Params["access_key"].(string)
		mpp.AWS.SecretKey = provider.Params["secret_key"].(string)
		mpp.AWS.Region = provider.Params["region"].(string)
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

type ManifestProviderAWS struct {
	AccessKey string
	SecretKey string
	Region    string
}
