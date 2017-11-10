package resources

type AwsLb struct {
	Name             string   `json:"name"`
	LoadBalancerType string   `json:"load_balancer_type"`
	IdleTimeout      string   `json:"idle_timeout"`
	SecurityGroups   []string `json:"security_groups"`
}

func NewAwsLb(name, securityGroupID string) AwsLb {
	return AwsLb{
		Name:             name,
		LoadBalancerType: "network",
		IdleTimeout:      "3600",
		SecurityGroups:   []string{securityGroupID},
	}
}
