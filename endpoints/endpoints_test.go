package endpoints

import (
	"reflect"
	"testing"

	"github.com/alekseiadamov/dendy/handlers"
)

func Test_unmarshal(t *testing.T) {
	tests := []struct {
		name           string
		endpointConfig []byte
		want           Endpoints
	}{
		{
			"EmptyConfig",
			[]byte(""),
			make(Endpoints),
		},
		{
			"OneEndpoint",
			[]byte(`hello:
  path: /
  method: GET
  handlerName: Hello`),
			Endpoints{
				"hello": {
					Path:        "/",
					Method:      "GET",
					HandlerName: "Hello",
					Handler:     nil,
				},
			},
		},
		{
			"AllEndpoints",
			[]byte(`hello:
  path: /
  method: GET
  handlerName: Hello
auth:
  path: /auth
  method: POST
  handlerName: Auth`),
			Endpoints{
				"hello": {
					Path:        "/",
					Method:      "GET",
					HandlerName: "Hello",
					Handler:     nil,
				},
				"auth": {
					Path:        "/auth",
					Method:      "POST",
					HandlerName: "Auth",
					Handler:     nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unmarshal(tt.endpointConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_setHandler(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *Endpoint
		want     *Endpoint
	}{
		// It's not possible to check func values for equality.
		// See reflect.DeepEqual documentation.
		{
			name: "HandlerDoesntExist",
			endpoint: &Endpoint{
				Path:        "/",
				Method:      "GET",
				HandlerName: "HelloThere",
				Handler:     nil,
			},
			want: &Endpoint{
				Path:        "/",
				Method:      "GET",
				HandlerName: "HelloThere",
				Handler:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.endpoint.Handler = handlers.Handlers[tt.endpoint.HandlerName]
			got := tt.endpoint
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
