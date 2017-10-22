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
		Params struct {
			Credentials string `yaml:"credentials"`
			Project     string `yaml:"project"`

			AccessKey string `yaml:"access_key"`
			SecretKey string `yaml:"secret_key"`

			SubscriptionID string `yaml:"subscription_id"`
			ClientID       string `yaml:"client_id"`
			ClientSecret   string `yaml:"client_secret"`
			TenantID       string `yaml:"tenant_id"`

			Region string `yaml:"region"`
			Zones  []Zone `yaml:"zones"`
		}
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
				Credentials: provider.Params.Credentials,
				Project:     provider.Params.Project,
				Region:      provider.Params.Region,
				Zones:       provider.Params.Zones,
			},
		}, nil
	case "aws":
		return Provider{
			Type: provider.Type,
			AWS: AWS{
				AccessKey: provider.Params.AccessKey,
				SecretKey: provider.Params.SecretKey,
				Region:    provider.Params.Region,
				Zones:     provider.Params.Zones,
			},
		}, nil
	case "azure":
		return Provider{
			Type: provider.Type,
			Azure: Azure{
				SubscriptionID: provider.Params.SubscriptionID,
				ClientID:       provider.Params.ClientID,
				ClientSecret:   provider.Params.ClientSecret,
				TenantID:       provider.Params.TenantID,
				Region:         provider.Params.Region,
			},
		}, nil
	default:
		return Provider{}, fmt.Errorf("unknown provider type: %q", provider.Type)
	}
}
