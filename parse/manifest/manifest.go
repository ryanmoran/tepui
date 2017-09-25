package manifest

type Manifest struct {
	Name          string         `yaml:"name"`
	Networks      []Network      `yaml:"networks"`
	LoadBalancers []LoadBalancer `yaml:"load_balancers"`
}
