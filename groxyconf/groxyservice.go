package groxyconf

type GroxyService interface {
	GetId() string
	GetPort() int
	GetHealthCheck() HealthCheck
}

type GroxyServiceStruct struct {
	Id          string            `yaml:"id,omitempty"`
	Port        int               `yaml:"port,omitempty"`
	HealthCheck HealthCheckStruct `yaml:"healthcheck,omitempty"`
}

func (this GroxyServiceStruct) GetPort() int {
	return this.Port
}

func (this GroxyServiceStruct) GetHealthCheck() HealthCheck {
	return this.HealthCheck
}

func (this GroxyServiceStruct) GetId() string {
	return this.Id
}
