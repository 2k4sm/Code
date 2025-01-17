package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
	}
	http.NotFound(w, r)

}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Myrouter!")
}
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
