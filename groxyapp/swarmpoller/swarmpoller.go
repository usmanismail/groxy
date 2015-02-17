package swarmpoller

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SwarmPoller interface {
	Start()
}

type SwarmPollerStruct struct {
	swarmUrl string
}

func CreateSwarmPoller(swarmUrl string) SwarmPoller {
	return SwarmPollerStruct{swarmUrl}
}

func (this SwarmPollerStruct) Start() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			resp, err := http.Get(this.swarmUrl)
			if err != nil {
				log.Printf("Error parsing docker swarm API response", err)
			} else {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Error parsing docker swarm API response", err)
				} else {
					log.Println(body)
				}
			}
		}
	}()
}
