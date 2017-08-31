package generate

import "encoding/json"

type TemplateProviderCollection struct {
	providers map[string]provider
}

func (tpc *TemplateProviderCollection) Add(p provider) {
	if tpc.providers == nil {
		tpc.providers = make(map[string]provider)
	}

	tpc.providers[p.ProviderName()] = p
}

func (tpc TemplateProviderCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(tpc.providers)
}
