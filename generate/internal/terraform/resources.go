package terraform

import "encoding/json"

type Resources []NamedResource

func (r Resources) MarshalJSON() ([]byte, error) {
	m := map[string]map[string]Resource{}

	for _, nr := range r {
		if m[nr.Type()] == nil {
			m[nr.Type()] = map[string]Resource{}
		}

		m[nr.Type()][nr.Name] = nr.Resource
	}

	return json.Marshal(m)
}
