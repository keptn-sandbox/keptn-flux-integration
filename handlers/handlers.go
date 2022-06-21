package handlers

import (
	"sync/atomic"

	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	r := mux.NewRouter()
	isReady := &atomic.Value{}
	isReady.Store(true)

	r.HandleFunc("/", home).Methods("POST")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}
