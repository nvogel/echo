package handlers

import (
   	"log"
    "net/http"
)

func readyz(w http.ResponseWriter, _ *http.Request) {
    log.Print("/readyz")
    w.WriteHeader(http.StatusOK)
}
