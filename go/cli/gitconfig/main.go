package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var ErrNotDirectoryFound = errors.New(".git not found, probably not the root repository.")

func main() {
	var path string
	flag.StringVar(&path, "p", ".", "Path for git base projects.")
	flag.Parse()

	gitconfig, err := walkGitConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gitconfig)
}

// find the gitconfig for a given repo and gets the url.
func walkGitConfig(repoPath string) (string, error) {
	repoPath, err := filepath.Abs(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	gitdir := path.Join(repoPath, ".git")
	config := path.Join(gitdir, "config")
	data, err := ioutil.ReadFile(config)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "file or directory") {
			err = ErrNotDirectoryFound
		}
		return "", err
	}
	re := regexp.MustCompile("url = (.*)")
	match := re.FindStringSubmatch(string(data))
	return match[1], nil
}
