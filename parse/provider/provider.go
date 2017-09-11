package provider

type Provider struct {
	Type  string
	GCP   GCP
	AWS   AWS
	Azure Azure
}

type GCP struct {
	Credentials string
	Project     string
	Region      string
}

type AWS struct {
	AccessKey string
	SecretKey string
	Region    string
}

type Azure struct {
	SubscriptionID string
	ClientID       string
	ClientSecret   string
	TenantID       string
	Region         string
}
