package handlers

import "net/http"

var Handlers = map[string]http.HandlerFunc{
	"Hello": Hello,
	"Auth":  Auth,
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func Auth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
	w.Write([]byte("You're not welcome here"))
}
