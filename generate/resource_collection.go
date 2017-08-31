package generate

import "encoding/json"

type TemplateResourceCollection struct {
	resources map[string]map[string]resource
}

func (trc *TemplateResourceCollection) Add(name string, r resource) {
	if trc.resources == nil {
		trc.resources = make(map[string]map[string]resource)
	}

	if trc.resources[r.ResourceType()] == nil {
		trc.resources[r.ResourceType()] = make(map[string]resource)
	}

	trc.resources[r.ResourceType()][name] = r
}

func (trc TemplateResourceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(trc.resources)
}
