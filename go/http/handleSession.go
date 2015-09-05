package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type CookieSession struct {
	Name string
	ID   string
}

func HandlerHelloSession(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	index, err := template.New("index").ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	cookie, err = r.Cookie("CookieExample")
	if err == http.ErrNoCookie {
		c := &CookieSession{Name: "CookieExample", ID: "2034"}
		cookie = &http.Cookie{
			Name:    c.Name,
			Value:   c.ID,
			Expires: time.Now().Add(time.Second * 60),
		}

		http.SetCookie(w, cookie)
		index.ExecuteTemplate(w, "index", "Set cookie for the first time, valid 1m. Refresh to see the content")
		return
	}

	if cookie != nil {
		index.ExecuteTemplate(w, "index", cookie)
		return
	}

}

func main() {
	fmt.Println("Listening on :9090")

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandlerHelloSession)
	http.ListenAndServe(":9090", mux)
}
