package main

import (
	"crypto/md5"
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	m bool
	s bool
)

func init() {
	flag.BoolVar(&m, "m", false, "hash md5")
	flag.BoolVar(&s, "s", false, "hash sha1")
	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "usage: checksum [options]  filename\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	filename, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		logrus.Fatal(err)
	}

	switch {
	case m:
		checksum := md5.Sum(filename)
		fmt.Printf("md5: %x\n", checksum)
	case s:
		checksum := sha1.Sum(filename)
		fmt.Printf("sha1: %x\n", checksum)
	}
}
