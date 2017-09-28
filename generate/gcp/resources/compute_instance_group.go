package resources

type GoogleComputeInstanceGroup struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
}

func NewGoogleComputeInstanceGroup(name, zone string) GoogleComputeInstanceGroup {
	return GoogleComputeInstanceGroup{
		Name: name,
		Zone: zone,
	}
}
