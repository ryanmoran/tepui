package resources

type GoogleComputeHealthCheck struct {
	Name           string                      `json:"name"`
	TCPHealthCheck GoogleComputeHealthCheckTCP `json:"tcp_health_check"`
}

type GoogleComputeHealthCheckTCP struct{}
