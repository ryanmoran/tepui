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
	case "azure":
		mpp.Azure.SubscriptionID = provider.Params["subscription_id"].(string)
		mpp.Azure.ClientID = provider.Params["client_id"].(string)
		mpp.Azure.ClientSecret = provider.Params["client_secret"].(string)
		mpp.Azure.TenantID = provider.Params["tenant_id"].(string)
		mpp.Azure.Region = provider.Params["region"].(string)
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
