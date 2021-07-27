// Package p contains an HTTP Cloud Function.
package p

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	log.Printf("%s, %s\n", b, err)
	w.WriteHeader(http.StatusOK)
}
