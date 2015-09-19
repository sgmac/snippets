package main

import (
	"fmt"
	"log"
	"net/http"

	"html/template"
)

const port = ":8080"

func HandlerRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.New("index").ParseFiles("templates/index.html")
		err = tmpl.ExecuteTemplate(w, "index", nil)
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
		fmt.Fprintf(w, "Welcome %s\n", name)
	}
}

func main() {
	log.Println("Listening on ", port)
	http.HandleFunc("/", HandlerRequest)
	http.ListenAndServe(port, nil)
}
