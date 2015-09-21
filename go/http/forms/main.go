package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"html/template"
)

const port = ":8080"

func HandlerRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/index.html",
			"templates/partials/header.html",
			"templates/partials/footer.html",
		))
		err := tmpl.ExecuteTemplate(w, "index", nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		name := r.Form.Get("name")
		envar := r.Form.Get("envar")
		value := os.Getenv(envar)
		fmt.Fprintf(w, "Welcome %s\n", name)
		fmt.Fprintf(w, "Your env  %s=%s\n", envar, value)
	}
}

func main() {
	log.Println("Listening on ", port)
	http.HandleFunc("/", HandlerRequest)
	http.ListenAndServe(port, nil)
}
