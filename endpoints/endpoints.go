package endpoints

import (
	"log"
	"net/http"
	"os"

	"github.com/alekseiadamov/dendy/handlers"
	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Path        string `yaml:"path"`
	Method      string `yaml:"method"`
	HandlerName string `yaml:"handlerName"`
	Handler     http.HandlerFunc
}

type Endpoints map[string]*Endpoint

func (endpoints Endpoints) setHandlers() {
	for _, endpoint := range endpoints {
		endpoint.Handler = handlers.Handlers[endpoint.HandlerName]
	}
}

func Read(configPath string) Endpoints {
	endpointConfig, fileReadingError := os.ReadFile(configPath)

	if fileReadingError != nil {
		log.Fatal(fileReadingError)
	}

	endpoints := unmarshal(endpointConfig)
	endpoints.setHandlers()
	return endpoints
}

func unmarshal(endpointConfig []byte) Endpoints {
	endpoints := make(Endpoints)

	unmarshallingError := yaml.Unmarshal(endpointConfig, &endpoints)

	if unmarshallingError != nil {
		log.Fatal(unmarshallingError)
	}

	return endpoints
}
