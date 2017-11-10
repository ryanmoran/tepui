package resources

type GoogleComputeInstanceGroup struct {
	Name    string `json:"name"`
	Zone    string `json:"zone"`
	Network string `json:"network"`
}

func NewGoogleComputeInstanceGroup(name, zone, network string) GoogleComputeInstanceGroup {
	return GoogleComputeInstanceGroup{
		Name:    name,
		Zone:    zone,
		Network: network,
	}
}
