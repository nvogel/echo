package handlers

import (
	"fmt"
    "log"
	"net/http"
)

// root is a simple HTTP handler function which writes a response.
func root(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Root")
	if err != nil {
        log.Printf("Could not write: %v", err)
    }
    log.Print("/root")
}
