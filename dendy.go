package dendy

import "github.com/alekseiadamov/dendy/endpoints"

// Reads an endpoint configuration file,
// creates corresponding endpoints and runs a server.
//
// Example:
//
//	package main
//
//	import (
//		"github.com/alekseiadamov/dendy"
//
//		"github.com/username/projectname/handlers"
//	)
//
//	func main() {
//		dendy.Serve("localhost:3333", "./example.yaml", handlers.Handlers)
//	}
func Serve(address string, configPath string, handlers endpoints.Handlers) {
	router := NewRouter()
	endpoints := endpoints.Read(configPath, handlers)

	router.CreateEndpoints(endpoints)
	router.Serve(address)
}
