# DENDY

(Which stands for **D**eclarative **EN**dpoint **D**efinition in **Y**aml.)

## How to use

1. Define API endpoints in YAML.

    ```yaml
    hello:
      path: /
      type: GET
      callbackName: Hello
    auth:
      path: /auth
      type: POST
      callbackName: Auth
    ```

2. Emplement endpoints callback functions in the `callbacks` package and put them in the `Callbacks` map.
   Unfortunately, Go doesn't allow to call package functions by their names without binding them somehow &mdash; be it a map or a type method. That's the reason for the requirement.

    ```go
    package callbacks

    import "net/http"

    var Callbacks = map[string]http.HandlerFunc{
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

3. Run server:

    ```go
    package main

    import "github.com/alekseiadamov/dendy"

    func main() {
        dendy.Serve("./example.yaml", "localhost:3333")
    }
    ```

4. Check response:

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
