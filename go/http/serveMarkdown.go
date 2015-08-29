package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/russross/blackfriday"
)

func ProcessMarkdownFile(filename string) []byte {
	fmt.Println("Requested /", filename)
	file, err := os.Open(filename)
	if err != nil {
		logrus.Fatal("error-", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Fatal("error-", err)
	}
	d := blackfriday.MarkdownBasic(data)
	return d
}

func HandlerMarkdown(w http.ResponseWriter, r *http.Request) {
	filenames := strings.Split(r.URL.Path, "/")
	if filenames[1] == "" {
		d := ProcessMarkdownFile("index.md")
		fmt.Fprintf(w, "%s", string(d))
		return
	}
	d := ProcessMarkdownFile(filenames[1])
	fmt.Fprintf(w, "%s", string(d))
}
func main() {

	fmt.Println("Listening port :9090")
	http.HandleFunc("/", HandlerMarkdown)
	http.ListenAndServe(":9090", nil)

}
