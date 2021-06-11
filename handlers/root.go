package handlers

import (
    "log"
	"net/http"
	"os"
	"html/template"
	"path"
)

type Profile struct {
	Color   string
  }

// root is a simple HTTP handler function which writes a response.
func root(w http.ResponseWriter, _ *http.Request) {
    color := "white"
    val, ok := os.LookupEnv("COLOR")
    if !ok {
        log.Print("COLOR env not set")
    } else {
        color = val
    }

	profile := Profile{color}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	if err := tmpl.Execute(w, profile)
	err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    log.Print("/")
}
