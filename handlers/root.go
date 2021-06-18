package handlers

import (
    "log"
	"net/http"
	"os"
	"html/template"
)

type Profile struct {
	Color   string
	Hostname string
	Kind string
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

	kind := "dog"
    val2, ok := os.LookupEnv("KIND")
    if !ok {
        log.Print("KIND env not set")
    } else {
        kind = val2
    }

    hostname, err := os.Hostname()

    if err != nil {
        log.Printf("Could not get hostname: %v", err)
        http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
        return
    }


	profile := Profile{color, hostname, kind}

	tmpl, err := template.New("foo").Parse(`
<!DOCTYPE html>
<html>
<body>

<p><em style="background-color:{{ .Color }}">?</em> <em>This is a {{ .Color }}</em> <em>{{ .Kind }}</em></p>
<div>
<p style="font-size: 0.5em">Generated on {{ .Hostname }} </p>
</div>
</body>
</html>
	`)
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
