package aws

import "encoding/json"

type Provider struct {
	AccessKey string
	SecretKey string
	Region    string
}

func (Provider) ProviderName() string { return "aws" }

func (p Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]map[string]string{
		"aws": map[string]string{
			"access_key": p.AccessKey,
			"secret_key": p.SecretKey,
			"region":     p.Region,
		},
	})
}
