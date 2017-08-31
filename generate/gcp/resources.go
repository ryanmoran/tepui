package gcp

import "encoding/json"

type ComputeNetworksCollection []ComputeNetwork

func (cnc ComputeNetworksCollection) MarshalJSON() ([]byte, error) {
	m := map[string]ComputeNetwork{}

	for _, computeNetwork := range cnc {
		m[computeNetwork.name] = computeNetwork
	}

	return json.Marshal(m)
}

type ComputeNetwork struct {
	name string

	Name string `json:"name"`
}

func (cn ComputeNetwork) ResourceType() string {
	return "google_compute_network"
}
