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
//		dendy.Serve("./example.yaml", "localhost:3333")
//	}
func Serve(configPath string, address string) {
	router := NewRouter()
	endpoints := endpoints.Read(configPath)

	router.CreateEndpoints(endpoints)
	router.Serve(address)
}
