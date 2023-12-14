package endpoints

import (
	"log"
	"net/http"
	"os"

	"github.com/alekseiadamov/dendy/callbacks"
	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Path         string `yaml:"path"`
	Method       string `yaml:"method"`
	CallbackName string `yaml:"callbackName"`
	Callback     http.HandlerFunc
}

type Endpoints map[string]*Endpoint

func (endpoints Endpoints) setCallbacks() {
	for _, endpoint := range endpoints {
		endpoint.Callback = callbacks.Callbacks[endpoint.CallbackName]
	}
}

func Read(configPath string) Endpoints {
	endpointConfig, fileReadingError := os.ReadFile(configPath)

	if fileReadingError != nil {
		log.Fatal(fileReadingError)
	}

	endpoints := unmarshal(endpointConfig)
	endpoints.setCallbacks()
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
