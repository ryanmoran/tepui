package resources

import (
	"strconv"

	"github.com/ryanmoran/tepui/generate/internal/terraform"
)

type GoogleComputeGlobalForwardingRule struct {
	Name      string `json:"name"`
	IPAddress string `json:"ip_address"`
	PortRange string `json:"port_range"`
	Target    string `json:"target"`
}

func NewGoogleComputeGlobalForwardingRule(name string, port int, address, target terraform.NamedResource) GoogleComputeGlobalForwardingRule {
	return GoogleComputeGlobalForwardingRule{
		Name:      name,
		PortRange: strconv.Itoa(port),
		IPAddress: address.Attribute("address"),
		Target:    target.SelfLink(),
	}
}
