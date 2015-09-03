package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

var config Config

type Config struct {
	Params params
	Runner runner
}
type params struct {
	Host    string
	Port    int
	Enabled bool
}
type runner struct {
	Commands []string
}

type dataHTTP struct{}

// Checks the command form the url is in the
// TOML config file.
func permitRunCommand(cmd string) bool {
	for _, c := range config.Runner.Commands {
		if strings.ToLower(c) == strings.ToLower(cmd) {
			return true
		}
	}
	return false
}
func (d dataHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPaths := strings.Split(r.URL.Path, "/")
	path := urlPaths[1]
	switch path {
	case "":
		// If empty this is root /
		fmt.Fprintf(w, "Index path\n")
		return
	case "user":
		var user string
		if urlPaths[2] == "" {
			user = "Nobody"
		}
		fmt.Fprintf(w, "Hello user %s\n", user)
		return
	case "run":
		run := urlPaths[2]

		if permitRunCommand(run) {
			cmd := exec.Command(run)
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "Command: %s\n", run)
			fmt.Fprintf(w, "Output: %s\n", string(stdout))
			return
		}
		fmt.Fprintf(w, "You do not have permission to run : %s\n", run)
		return
	default:
		fmt.Fprintf(w, "This url wasn't prepared")

	}
}

func main() {

	var serv dataHTTP
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}

	listen := config.Params.Host + ":" + strconv.Itoa(config.Params.Port)
	fmt.Println("Listening on ", listen)
	http.ListenAndServe(listen, serv)
}
