package manifest

type ManifestNetwork struct {
	Name    string           `yaml:"name"`
	CIDR    string           `yaml:"cidr"`
	Subnets []ManifestSubnet `yaml:"subnets"`
}
