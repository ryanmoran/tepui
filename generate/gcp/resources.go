package gcp

import "encoding/json"

type Resources []NamedResource

func (r Resources) MarshalJSON() ([]byte, error) {
	m := map[string]Resource{}

	for _, nr := range r {
		m[nr.Name] = nr.Resource
	}

	return json.Marshal(m)
}

type NamedResource struct {
	Name     string
	Resource Resource
}

type Resource interface{}

type ComputeNetwork struct {
	Name string `json:"name"`
}

func (cn ComputeNetwork) ResourceType() string {
	return "google_compute_network"
}
