package main

import (
	"fmt"
	"net/http"
)

// Index Just return something
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
