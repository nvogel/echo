package handlers

import (
	"fmt"
    "log"
	"net/http"
)

// home is a simple HTTP handler function which writes a response.
func root(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "/help")
    log.Print("/root")
}
