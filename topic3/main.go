// http server draws a surface.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("http server is running")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("got new request: %v %v %v\n", r.Method, r.URL.String, r.RemoteAddr)
	w.Header().Set("Content-Type", "image/svg+xml")
	surface := Surface()

	fmt.Fprintln(w, surface)
}
