package manifest

type LoadBalancer struct {
	Name  string   `yaml:"name"`
	Ports []int    `yaml:"ports"`
	Zones []string `yaml:"zones"`
}
