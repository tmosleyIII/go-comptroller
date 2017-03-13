package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var err error

// Index Just return something
func Index(w http.ResponseWriter, r *http.Request) {
	span := traceClient.SpanFromRequest(r)
	defer span.Finish()
	fmt.Fprintf(w, "Hello World!")
}

// healthHandler return OK if application is running
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

// ReadinessHandler checks status of dependencies
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	ok := true
	errMsg := ""
	if db != nil {
		_, err := db.Query("SELECT 1 from foo;")
		if err != nil {
			ok = false
			errMsg += "Database not ok."
			log.Println(err)
		}
	}
	if db == nil {
		ok = false
		errMsg += "Database not ok."
	}

	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, errMsg, http.StatusServiceUnavailable)
	}
}
