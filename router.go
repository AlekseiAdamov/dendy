package dendy

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/alekseiadamov/dendy/endpoints"
	"github.com/go-chi/chi/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const defaultAddress = "localhost:3333"

type Router struct {
	*chi.Mux
}

func NewRouter() Router {
	return Router{chi.NewRouter()}
}

func (router Router) CreateEndpoints(endpoints endpoints.Endpoints) {
	value := reflect.ValueOf
	caser := cases.Title(language.English)

	for _, endpoint := range endpoints {
		params := []reflect.Value{value(endpoint.Path), value(endpoint.Handler)}
		method := caser.String(strings.ToLower(endpoint.Method))
		routerMethod := value(router).MethodByName(method)
		routerMethod.Call(params)
	}
}

func (router Router) Serve(address string) error {
	if len(strings.TrimSpace(address)) == 0 {
		address = defaultAddress
	}
	return http.ListenAndServe(address, router)
}
