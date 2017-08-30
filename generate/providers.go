package generate

type TemplateProviderGoogle struct {
	Credentials string `json:"credentials"`
	Project     string `json:"project"`
	Region      string `json:"region"`
}

type TemplateProviderAWS struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

type TemplateProviderAzure struct {
	SubscriptionID string `json:"subscription_id"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	TenantID       string `json:"tenant_id"`
}

type provider interface {
	_provider()
}

func (tpg TemplateProviderGoogle) _provider() {}
func (tpa TemplateProviderAWS) _provider()    {}
func (tpa TemplateProviderAzure) _provider()  {}
