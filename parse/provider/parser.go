package provider

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func (p Parser) Parse(path string) (Provider, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return Provider{}, err
	}

	var provider struct {
		Type   string `yaml:"type"`
		Params map[string]string
	}

	err = yaml.Unmarshal(contents, &provider)
	if err != nil {
		return Provider{}, err
	}

	switch provider.Type {
	case "gcp":
		return Provider{
			Type: provider.Type,
			GCP: GCP{
				Credentials: provider.Params["credentials"],
				Project:     provider.Params["project"],
				Region:      provider.Params["region"],
			},
		}, nil
	case "aws":
		return Provider{
			Type: provider.Type,
			AWS: AWS{
				AccessKey: provider.Params["access_key"],
				SecretKey: provider.Params["secret_key"],
				Region:    provider.Params["region"],
			},
		}, nil
	case "azure":
		return Provider{
			Type: provider.Type,
			Azure: Azure{
				SubscriptionID: provider.Params["subscription_id"],
				ClientID:       provider.Params["client_id"],
				ClientSecret:   provider.Params["client_secret"],
				TenantID:       provider.Params["tenant_id"],
				Region:         provider.Params["region"],
			},
		}, nil
	default:
		return Provider{}, fmt.Errorf("unknown provider type: %q", provider.Type)
	}
}
