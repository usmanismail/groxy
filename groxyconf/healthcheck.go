package groxyconf

type HealthCheck interface {
	GetUri() string
	GetUnhealthyThreshold() string
	GetHealthyThreshold() string
}

type HealthCheckStruct struct {
	URI                string `yaml:"uri,omitempty"`
	UnhealthyThreshold string `yaml:"unhealthy_threshold,omitempty"`
	HealthyThreshold   string `yaml:"healthy_threshold,omitempty"`
}

func (this HealthCheckStruct) GetUri() string {
	return this.URI
}

func (this HealthCheckStruct) GetUnhealthyThreshold() string {
	return this.UnhealthyThreshold
}

func (this HealthCheckStruct) GetHealthyThreshold() string {
	return this.HealthyThreshold
}
