package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home").ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "home", nil)
	return
}

func HandlerListEnv(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("list").ParseFiles("templates/list.html")
	if err != nil {
		log.Fatal(err)
	}
	envars := os.Environ()
	vars := make(map[string]string)
	for _, envarvalue := range envars {
		envar := strings.Split(envarvalue, "=")[0]
		value := strings.Split(envarvalue, "=")[1]
		vars[envar] = value
	}

	t.ExecuteTemplate(w, "list", vars)

	return
}

func main() {
	fmt.Println("Listening in :9090")
	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/listenv", HandlerListEnv)

	// Serves static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9090", nil)
}
