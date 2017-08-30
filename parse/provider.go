package parse

type Provider struct {
	Type  string
	GCP   ProviderGCP
	AWS   ProviderAWS
	Azure ProviderAzure
}

type ProviderGCP struct {
	Credentials string
	Project     string
	Region      string
}

type ProviderAWS struct {
	AccessKey string
	SecretKey string
	Region    string
}

type ProviderAzure struct {
	SubscriptionID string
	ClientID       string
	ClientSecret   string
	TenantID       string
	Region         string
}
