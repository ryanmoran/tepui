package resources

type AwsLbTargetGroup struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	VPCID    string `json:"vpc_id"`
}

func NewAwsLbTargetGroup(name, vpcID string, port int) AwsLbTargetGroup {
	return AwsLbTargetGroup{
		Name:     name,
		Port:     port,
		Protocol: "http",
		VPCID:    vpcID,
	}
}
