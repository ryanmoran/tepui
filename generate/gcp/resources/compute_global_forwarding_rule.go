package resources

type GoogleComputeGlobalForwardingRule struct {
	Name      string `json:"name"`
	IPAddress string `json:"ip_address"`
	PortRange string `json:"port_range"`
	Target    string `json:"target"`
}
