package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rtfreedman/SpellTracker/api/routes"
)

// RunAPI starts the API
func RunAPI(port int, external bool) {
	// if we aren't running an internal instance we want to bind 0.0.0.0 so the API is
	// externally accessible
	loc := "0.0.0.0:%d"
	if !external {
		loc = "127.0.0.1:%d"
	}
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(loc, port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	routes.SetupSpells(r)
	log.Fatal(srv.ListenAndServe())
}
