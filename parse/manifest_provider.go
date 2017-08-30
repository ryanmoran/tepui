package parse

import "fmt"

type ManifestProvider struct {
	Type  string
	GCP   ManifestProviderGCP
	AWS   ManifestProviderAWS
	Azure ManifestProviderAzure
}

func (mpp *ManifestProvider) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var provider struct {
		Type   string `yaml:"type"`
		Params map[string]string
	}

	err := unmarshal(&provider)
	if err != nil {
		return err
	}

	mpp.Type = provider.Type

	switch provider.Type {
	case "gcp":
		mpp.GCP.Credentials = provider.Params["credentials"]
		mpp.GCP.Project = provider.Params["project"]
		mpp.GCP.Region = provider.Params["region"]
	case "aws":
		mpp.AWS.AccessKey = provider.Params["access_key"]
		mpp.AWS.SecretKey = provider.Params["secret_key"]
		mpp.AWS.Region = provider.Params["region"]
	case "azure":
		mpp.Azure.SubscriptionID = provider.Params["subscription_id"]
		mpp.Azure.ClientID = provider.Params["client_id"]
		mpp.Azure.ClientSecret = provider.Params["client_secret"]
		mpp.Azure.TenantID = provider.Params["tenant_id"]
		mpp.Azure.Region = provider.Params["region"]
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

type ManifestProviderAzure struct {
	SubscriptionID string
	ClientID       string
	ClientSecret   string
	TenantID       string
	Region         string
}
