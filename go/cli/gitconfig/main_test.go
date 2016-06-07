package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testDirectory = "/tmp/testing"

// var testDirectory = "/tmp/dotfiles"

var stubGitConfig = `
[remote "origin"]
        url = git@gitlab.org:stubuser/testing.git
	fetch = +refs/heads/*:refs/remotes/origin/*
`

func TestWalkGitConfig(t *testing.T) {
	defer os.RemoveAll(testDirectory)

	err := createDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	url, err := walkGitConfig(testDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	if url != "git@gitlab.org:stubuser/testing.git" {
		t.Errorf("repository url did not match, got %s\n", url)
	}
}

// create a fake directory with a fake git config file.
func createDir() error {
	dirPath := filepath.Join(testDirectory, ".git")
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s", err.Error())
		os.Exit(1)
	}
	filePath := filepath.Join(dirPath, "config")
	ioutil.WriteFile(filePath, []byte(stubGitConfig), 0644)

	return nil
}
