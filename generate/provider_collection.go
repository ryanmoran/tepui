package generate

import "encoding/json"

type TemplateProviderCollection struct {
	providers map[string]provider
}

func (tpc *TemplateProviderCollection) Add(p provider) {
	if tpc.providers == nil {
		tpc.providers = make(map[string]provider)
	}

	switch p.(type) {
	case TemplateProviderGoogle:
		tpc.providers["google"] = p
	case TemplateProviderAWS:
		tpc.providers["aws"] = p
	}
}

func (tpc TemplateProviderCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(tpc.providers)
}
