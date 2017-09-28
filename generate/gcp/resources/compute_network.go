package resources

type GoogleComputeNetwork struct {
	Name string `json:"name"`
}

func NewGoogleComputeNetwork(name string) GoogleComputeNetwork {
	return GoogleComputeNetwork{
		Name: name,
	}
}
