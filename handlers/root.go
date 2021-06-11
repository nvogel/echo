package handlers

import (
    "log"
	"net/http"
	"os"
	"html/template"
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

	tmpl, err := template.New("foo").Parse(`
<!DOCTYPE html>
<html>
<body style="background-color:{{ .Color }};">

<h1>{{ .Color }}</h1>
<div>
    <p>
    <a href="./whoami">Who am i</a>

    </p>
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
