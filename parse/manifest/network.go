package manifest

type Network struct {
	Name          string         `yaml:"name"`
	CIDR          string         `yaml:"cidr"`
	Subnets       []Subnet       `yaml:"subnets"`
	LoadBalancers []LoadBalancer `yaml:"load_balancers"`
}
