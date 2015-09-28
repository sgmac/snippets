package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var (
	user  = flag.String("u", "", "")
	usage = `usage: qpass [options]
 Options:
   -u Username
`
)

func main() {
	pswd := make(map[string]map[string]string)
	var wg sync.WaitGroup

	response := make(chan interface{})
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	data, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error read file %s\n", err)
	}
	data2str := string(data)
	lines := strings.Split(data2str, "\n")

	wg.Add(1)
	go func() {
		for _, l := range lines {
			if strings.HasPrefix(l, "#") {
				continue
			}
			if l == "" {
				break
			}
			fields := strings.Split(l, ":")
			username := fields[0]
			pswd[username] = make(map[string]string)
			pswd[username]["uid"] = fields[2]
			pswd[username]["gid"] = fields[3]
			pswd[username]["gecos"] = fields[4]
			pswd[username]["home"] = fields[5]
			pswd[username]["shell"] = fields[6]
		}
		wg.Done()
		response <- pswd
	}()

	wg.Wait()
	m := <-response
	v := m.(map[string]map[string]string)

	fmt.Printf("%s\n", v[*user])
}
