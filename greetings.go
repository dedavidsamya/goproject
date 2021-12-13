package main

import (
	"fmt"
	"net/http"
)

// Hello returns a greeting for the named person.
func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
func hello(w http.ResponseWriter, req *http.Request) {
	name, _ := req.URL.Query()["name"]
	fmt.Fprintf(w, string(name[0]))

}
