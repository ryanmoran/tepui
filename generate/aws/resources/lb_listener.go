package resources

import "strconv"

type AwsLbListener struct {
	LoadBalancerARN string `json:"load_balancer_arn"`
	Port            string `json:"port"`
	DefaultAction   struct {
		TargetGroupARN string `json:"target_group_arn"`
		Type           string `json:"type"`
	} `json:"default_action"`
}

func NewAwsLbListener(loadBalancerARN, targetGroupARN string, port int) AwsLbListener {
	listener := AwsLbListener{
		LoadBalancerARN: loadBalancerARN,
		Port:            strconv.Itoa(port),
	}
	listener.DefaultAction.TargetGroupARN = targetGroupARN
	listener.DefaultAction.Type = "forward"

	return listener
}
