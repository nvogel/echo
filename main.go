package main

import (
	"fmt"
	"net/http"
    "log"
    "github.com/nvogel/echo/version"
)

func main() {

    log.Print("Starting the service...")
    log.Printf("Infos commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)

	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello! Your request was processed.")
	},
	)

    log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
