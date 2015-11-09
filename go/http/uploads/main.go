package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alecthomas/template"
)

const server = ":8000"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		temp, err := template.New("upload").ParseGlob("templates/*")
		if err != nil {
			fmt.Fprintf(w, "Error: ", err)
			return
		}
		temp.ExecuteTemplate(w, "upload", nil)

	case "POST":
		fmt.Fprintf(w, "Not implemented")
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/upload", 301)
	})
	http.HandleFunc("/upload", uploadHandler)

	fmt.Fprintf(os.Stdout, "Listening on %s", server)
	http.ListenAndServe(server, nil)
}
