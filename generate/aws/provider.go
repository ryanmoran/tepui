package aws

import "encoding/json"

type Provider struct {
	AccessKey string
	SecretKey string
	Region    string
}

func (Provider) ProviderName() string { return "aws" }

func (p Provider) MarshalJSON() ([]byte, error) {
	var t struct {
		AWS struct {
			AccessKey string `json:"access_key"`
			SecretKey string `json:"secret_key"`
			Region    string `json:"region"`
		} `json:"aws"`
	}

	t.AWS.AccessKey = p.AccessKey
	t.AWS.SecretKey = p.SecretKey
	t.AWS.Region = p.Region

	return json.Marshal(t)
}
