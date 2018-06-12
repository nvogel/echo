package handlers

import (
   	"log"
    "encoding/json"
    "net/http"
    "os"
    "github.com/nvogel/echo/version"
)

func whoami(w http.ResponseWriter, req *http.Request) {

    hostname, err := os.Hostname()

    if err != nil {
        log.Printf("Could not get hostname: %v", err)
        http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
        return
    }


    xForwardedFor := req.Header.Get("X-Forwarded-For")
    clientIp := req.RemoteAddr

    info := struct {
        BuildTime string `json:"buildTime"`
        Commit    string `json:"commit"`
        Release   string `json:"release"`
        Hostname  string `json:"hostname"`
        XForwardedFor string `json:"x-forwarded-for"`
        ClientIp      string `json:"client"`
    }{
        version.BuildTime, version.Commit, version.Release, hostname, xForwardedFor, clientIp,
    }

    body, err := json.Marshal(info)
    if err != nil {
        log.Printf("Could not encode info data: %v", err)
        http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _, err = w.Write(body)
    if err != nil {
        log.Printf("Could not write: %v", err)
    }

    log.Print("/whoami")
}
