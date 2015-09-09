package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/libgit2/git2go"
)

func main() {

	fmt.Println("Reading dotfiles repository")
	gitmain := os.Getenv("GIT")

	dotfiles := path.Join(gitmain, "dotfiles")
	repo, err := git.OpenRepository(dotfiles)
	if err != nil {
		logrus.Fatal(err)
	}

	head, err := repo.Head()
	if err != nil {
		logrus.Fatal(err)
	}

	// Last commit form HEAD, do not assume
	// HEAD points to master.
	lastCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
		logrus.Fatal(err)
	}
	author := lastCommit.Author()
	fmt.Printf("ID: %s\n", lastCommit.Id().String())
	fmt.Printf("Author: %s <%s> Date: %s\n ", author.Name, author.Email, author.When.Format(time.UnixDate))
	time.Now()
}
