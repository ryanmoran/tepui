package resources

type GoogleComputeHealthCheck struct {
	Name           string                      `json:"name"`
	TCPHealthCheck GoogleComputeHealthCheckTCP `json:"tcp_health_check"`
}

func NewGoogleComputeHealthCheck(name string) GoogleComputeHealthCheck {
	return GoogleComputeHealthCheck{
		Name:           name,
		TCPHealthCheck: GoogleComputeHealthCheckTCP{},
	}
}

type GoogleComputeHealthCheckTCP struct{}
