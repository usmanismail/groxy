package groxyconf

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type GroxyConfig interface {
	GetSwarmHost() string
	GetSwarmVersion() string
	GetServices() <-chan GroxyService
}

type GroxyConfiguration struct {
	SwarmHost    string               `yaml:"Swarm_host,omitempty"`
	SwarmVersion string               `yaml:"Swarm_version,omitempty"`
	Services     []GroxyServiceStruct `yaml:"Services,omitempty"`
}

func parseConfig(configFile *os.File) (GroxyConfig, error) {
	config := &GroxyConfiguration{}

	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}

	if config.SwarmHost == "" {
		return nil, errors.New("Missing Swarm Host Parameter")
	}

	if config.SwarmVersion == "" {
		return nil, errors.New("Missing Swarm Version Parameter")
	}

	log.Printf("Using %s to connect to Docker Swarm", config.SwarmHost)
	for _, value := range config.Services {
		log.Printf("Configuring Service %s, with health check on (%s)", value.GetId(), value.GetHealthCheck().GetUri())
		log.Printf("Healthy Threshold %s, Unhealthy Threshold %s",
			value.GetHealthCheck().GetHealthyThreshold(), value.GetHealthCheck().GetUnhealthyThreshold())
	}
	return config, nil
}

func (this GroxyConfiguration) GetSwarmHost() string {
	return this.SwarmHost
}

func (this GroxyConfiguration) GetSwarmVersion() string {
	return this.SwarmVersion
}

func (this GroxyConfiguration) GetServices() <-chan GroxyService {
	c := make(chan GroxyService)
	go func() {
		for _, service := range this.Services {
			c <- service
		}
		close(c)
	}()
	return c
}
