package resources

type GoogleComputeGlobalAddress struct {
	Name string `json:"name"`
}

func NewGoogleComputeGlobalAddress(name string) GoogleComputeGlobalAddress {
	return GoogleComputeGlobalAddress{
		Name: name,
	}
}
