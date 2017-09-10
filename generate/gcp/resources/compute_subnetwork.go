package resources

type GoogleComputeSubnetwork struct {
	Name        string `json:"name"`
	IPCIDRRange string `json:"ip_cidr_range"`
	Network     string `json:"network"`
}
