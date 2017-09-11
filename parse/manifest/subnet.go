package manifest

type Subnet struct {
	Name string `yaml:"name"`
	CIDR string `yaml:"cidr"`
}
