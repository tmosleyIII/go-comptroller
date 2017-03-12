package main

import (
	"fmt"
	"net/http"
)

// Index Just return something
func Index(w http.ResponseWriter, r *http.Request) {
	span := traceClient.SpanFromRequest(r)
	defer span.Finish()
	fmt.Fprintf(w, "Hello World!")
}
