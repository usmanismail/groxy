package groxyconf

import (
	"flag"
	"log"
	"os"
)

func ParseConfig() GroxyConfig {

	var config_file string
	flag.StringVar(&config_file, "config", "/etc/groxy/services.conf", "The path to the services configuration")
	flag.Parse()

	log.Printf("Reading services configuration from %s", config_file)
	configReader, err := os.Open(config_file)
	defer configReader.Close()

	if err != nil {
		log.Panicf("%s: %v", "Unable to read services config file", err)
	}

	config, err := parseConfig(configReader)

	if err != nil {
		log.Panicf("%s: %v", "Unable to parse service config", err)
	}

	return config
}
