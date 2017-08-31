package aws

type Provider struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

func (Provider) ProviderName() string { return "aws" }
