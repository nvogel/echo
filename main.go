package main

import (
	"net/http"
    "log"
    "github.com/nvogel/echo/version"
    "github.com/nvogel/echo/handlers"
    "os"
)

func main() {

    port := "8080"
    val, ok := os.LookupEnv("PORT")
    if !ok {
        log.Print("Port not set")
    } else {
        port = val
    }

    log.Print("Starting the service...")
    log.Printf("Infos commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)

    router := handlers.Router()

    log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":" + port, router))
}
