# DENDY

(Which stands for **D**eclarative **EN**dpoint **D**efinition in **Y**aml.)

## How to use

1. Define API endpoints in YAML.

    ```yaml
    hello:
      path: /
      type: GET
      handlerName: Hello
    auth:
      path: /auth
      type: POST
      handlerName: Auth
    ```

2. Implement endpoints handler functions and put them in the `map[string]http.HandlerFunc`.
   Unfortunately, Go doesn't allow to call package functions by their names without binding them somehow &mdash; be it a map or a type method. That's the reason for the requirement.

    ```go
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
    ```

3. Run server, passing previously created `map[string]http.HandlerFunc` as an argument.

    ```go
    package main

    import (
        "github.com/alekseiadamov/dendy"

        "github.com/username/projectname/handlers"
    )

    func main() {
        dendy.Serve("localhost:3333", "./example.yaml", handlers.Handlers)
    }
    ```

4. Check response.

    ```http
    GET http://localhost:3333

    HTTP/1.1 200 OK
    Date: Thu, 14 Dec 2023 16:27:09 GMT
    Content-Length: 12
    Content-Type: text/plain; charset=utf-8
    Connection: close

    Hello World!
    ```

    ```http
    POST http://localhost:3333/auth

    HTTP/1.1 403 Forbidden
    Date: Thu, 14 Dec 2023 16:28:59 GMT
    Content-Length: 23
    Content-Type: text/plain; charset=utf-8
    Connection: close

    You're not welcome here
    ```

## Dependencies

See [go.mod](./go.mod).
