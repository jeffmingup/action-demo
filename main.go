package main

import (
	"fmt"
	"net/http"
)

// define a type for the response
type Hello struct{}

// let that type implement the ServeHTTP method (defined in interface
// http.Handler)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello111!")
}

func main() {
	var h Hello
	http.ListenAndServe(":4000", h)
}
