package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func helloHandlerParam(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Fprintf(w, "Hello %s\n", v["name"])
}
func main() {
	server := mux.NewRouter()
	server.Path("/").Methods("GET").HandlerFunc(helloHandler)
	server.Path("/{name}").Methods("GET").HandlerFunc(helloHandlerParam)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", server)
}
