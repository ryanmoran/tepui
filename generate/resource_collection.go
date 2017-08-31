package generate

import "encoding/json"

type TemplateResourceCollection struct {
	resources map[string]map[string]Resource
}

func (trc *TemplateResourceCollection) Add(name string, r Resource) {
	if trc.resources == nil {
		trc.resources = make(map[string]map[string]Resource)
	}

	if trc.resources[r.ResourceType()] == nil {
		trc.resources[r.ResourceType()] = make(map[string]Resource)
	}

	trc.resources[r.ResourceType()][name] = r
}

func (trc TemplateResourceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(trc.resources)
}
