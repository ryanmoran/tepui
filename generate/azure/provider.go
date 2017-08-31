package azure

import "encoding/json"

type Provider struct {
	SubscriptionID string
	ClientID       string
	ClientSecret   string
	TenantID       string
}

func (p Provider) ProviderName() string { return "azurerm" }

func (p Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]map[string]string{
		"azurerm": map[string]string{
			"subscription_id": p.SubscriptionID,
			"client_id":       p.ClientID,
			"client_secret":   p.ClientSecret,
			"tenant_id":       p.TenantID,
		},
	})
}
