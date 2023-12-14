package endpoints

import (
	"reflect"
	"testing"

	"github.com/alekseiadamov/dendy/callbacks"
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
  callbackName: Hello`),
			Endpoints{
				"hello": {
					Path:         "/",
					Method:       "GET",
					CallbackName: "Hello",
					Callback:     nil,
				},
			},
		},
		{
			"AllEndpoints",
			[]byte(`hello:
  path: /
  method: GET
  callbackName: Hello
auth:
  path: /auth
  method: POST
  callbackName: Auth`),
			Endpoints{
				"hello": {
					Path:         "/",
					Method:       "GET",
					CallbackName: "Hello",
					Callback:     nil,
				},
				"auth": {
					Path:         "/auth",
					Method:       "POST",
					CallbackName: "Auth",
					Callback:     nil,
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

func TestEndpoint_setCallback(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *Endpoint
		want     *Endpoint
	}{
		// It's not possible to check func values for equality.
		// See reflect.DeepEqual documentation.
		{
			name: "CallbackDoesntExist",
			endpoint: &Endpoint{
				Path:         "/",
				Method:       "GET",
				CallbackName: "HelloThere",
				Callback:     nil,
			},
			want: &Endpoint{
				Path:         "/",
				Method:       "GET",
				CallbackName: "HelloThere",
				Callback:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.endpoint.Callback = callbacks.Callbacks[tt.endpoint.CallbackName]
			got := tt.endpoint
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
