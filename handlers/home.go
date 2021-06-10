package handlers

import (
	"fmt"
    "log"
	"net/http"
)

// home is a simple HTTP handler function which writes a response.
func home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Hello! Your request was processed.")
	if err != nil {
        log.Printf("Could not write: %v", err)
    }
	log.Print("/home")
}
