package handlers

import (
   	"log"
    "net/http"
)

func healthz(w http.ResponseWriter, _ *http.Request) {
    log.Print("/healthz")
    w.WriteHeader(http.StatusOK)
}
