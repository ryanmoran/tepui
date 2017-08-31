package gcp

type ComputeNetwork struct {
	Name string `json:"name"`
}

func (cn ComputeNetwork) ResourceType() string {
	return "google_compute_network"
}
