package resources

import "github.com/ryanmoran/tepui/generate/internal/terraform"

type GoogleComputeSubnetwork struct {
	Name        string `json:"name"`
	IPCIDRRange string `json:"ip_cidr_range"`
	Network     string `json:"network"`
}

func NewGoogleComputeSubnetwork(name, cidr string, network terraform.NamedResource) GoogleComputeSubnetwork {
	return GoogleComputeSubnetwork{
		Name:        name,
		IPCIDRRange: cidr,
		Network:     network.SelfLink(),
	}
}
