package endpoints

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Path        string `yaml:"path"`
	Method      string `yaml:"method"`
	HandlerName string `yaml:"handlerName"`
	Handler     http.HandlerFunc
}

type Endpoints map[string]*Endpoint
type Handlers map[string]http.HandlerFunc

func (endpoints Endpoints) setHandlers(handlers Handlers) {
	for _, endpoint := range endpoints {
		endpoint.Handler = handlers[endpoint.HandlerName]
	}
}

func Read(configPath string, handlers Handlers) Endpoints {
	endpointConfig, fileReadingError := os.ReadFile(configPath)

	if fileReadingError != nil {
		log.Fatal(fileReadingError)
	}

	endpoints := unmarshal(endpointConfig)
	endpoints.setHandlers(handlers)
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
