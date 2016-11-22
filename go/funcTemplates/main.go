package main

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

// Uptime gets the uptime on the current host.
func Uptime() string {
	var out bytes.Buffer
	cmd := exec.Command("uptime")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

const tmpl = `
Uptime: {{ uptime }}
`

func customHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index").Funcs(template.FuncMap{"uptime": Uptime}).Parse(tmpl))
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	port := flag.Int("p", 8080, "Listen port.")
	flag.Parse()

	http.HandleFunc("/", customHandler)
	strPort := strconv.Itoa(*port)

	log.Printf("Listening on port %d\n", *port)
	http.ListenAndServe(":"+strPort, nil)
}
