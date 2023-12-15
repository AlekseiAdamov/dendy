package dendy

import "github.com/alekseiadamov/dendy/endpoints"

// Reads an endpoint configuration file,
// creates corresponding endpoints and runs a server.
//
// Example:
//
//	package main
//
//	import "github.com/alekseiadamov/dendy"
//
//	func main() {
//		dendy.Serve("localhost:3333", "./example.yaml")
//	}
func Serve(address string, configPath string) {
	router := NewRouter()
	endpoints := endpoints.Read(configPath)

	router.CreateEndpoints(endpoints)
	router.Serve(address)
}
