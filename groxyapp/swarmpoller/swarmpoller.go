package swarmpoller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/usmanismail/groxy/groxyapp/swarmpoller/meta"
)

type SwarmPoller interface {
	Start()
}

type SwarmPollerStruct struct {
	swarmUrl     string
	swarmVersion string
}

func CreateSwarmPoller(swarmUrl string, swarmVersion string) SwarmPoller {
	return SwarmPollerStruct{swarmUrl, swarmVersion}
}

func (this SwarmPollerStruct) Start() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			resp, err := http.Get(fmt.Sprintf("%s/%s/containers/json", this.swarmUrl, this.swarmVersion))
			if err != nil {
				log.Printf("Error parsing docker swarm API response", err)
			} else {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Error parsing docker swarm API response", err)
				} else {
					var containers []meta.ContainerMetaStruct
					err := json.Unmarshal(body, &containers)
					if err != nil {
						log.Printf("Error parsing docker swarm API response", err)
					} else {
						for _, container := range containers {
							switch {
							case strings.Contains(container.Status, "Up") == false:
								log.Printf("Container %s is not Up", container.Id)
							case len(container.Ports) == 0:
								log.Printf("Container %s does not have any ports", container.Id)
							default:
								log.Printf("Found Container: %s", container.Id)
							}
						}
					}
				}
			}
		}
	}()
}
