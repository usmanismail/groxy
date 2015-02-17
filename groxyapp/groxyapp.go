package groxyapp

import (
	"log"
	"sync"

	"github.com/usmanismail/groxy/groxyapp/swarmpoller"
	"github.com/usmanismail/groxy/groxyconf"
)

type GroxyApp interface {
	Start() error
}
type GroxyApplication struct {
	config groxyconf.GroxyConfig
}

func New() (GroxyApp, error) {
	config := groxyconf.ParseConfig()
	groxyapp_ := &GroxyApplication{config}
	return groxyapp_, nil
}

func (this GroxyApplication) Start() error {

	var wg sync.WaitGroup

	log.Println("Starting DockerSwarm Poller")
	_swarmpoller := swarmpoller.CreateSwarmPoller(this.config.GetSwarmHost(), this.config.GetSwarmVersion())
	_swarmpoller.Start()
	wg.Add(1)

	log.Println("Starting Services")
	services := this.config.GetServices()
	for service := range services {
		log.Printf("Service %s, Port %d, Health Check: %v",
			service.GetId(), service.GetPort(), service.GetHealthCheck().GetUri())
		wg.Add(1)
	}

	wg.Wait()
	return nil
}
