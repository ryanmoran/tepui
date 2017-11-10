package resources

import "strconv"

type AwsSecurityGroupRule struct {
	Type            string `json:"type"`
	Protocol        string `json:"protocol"`
	FromPort        string `json:"from_port"`
	ToPort          string `json:"to_port"`
	SecurityGroupID string `json:"security_group_id"`
}

func NewAwsSecurityGroupRule(securityGroupID string, port int) AwsSecurityGroupRule {
	return AwsSecurityGroupRule{
		Type:            "ingress",
		Protocol:        "tcp",
		FromPort:        strconv.Itoa(port),
		ToPort:          strconv.Itoa(port),
		SecurityGroupID: securityGroupID,
	}
}
