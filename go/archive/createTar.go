package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	files       []string
	filesstring string
	tarfile     string
)

func init() {
	flag.StringVar(&tarfile, "t", "content.tar", "Tar filename.")
	flag.StringVar(&filesstring, "f", "", "Files to add to the tarfile, separated by comma.")

}

func helpHeader() {
	fmt.Println("usage: ctar [options]")
}

func main() {
	flag.Parse()
	if flag.NFlag() < 1 {
		helpHeader()
		flag.PrintDefaults()
		os.Exit(1)
	}

	t, err := os.Create(tarfile)
	if err != nil {
		log.Fatal(err)
	}

	tw := tar.NewWriter(t)
	if !strings.Contains(filesstring, ",") {
		files = append(files, filesstring)
	} else {
		files = strings.Split(filesstring, ",")
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		// Set current time for each file, otherwise sets date
		// for each file on the archive to 1 Jun 1970.
		h := tar.Header{
			ModTime:    time.Now(),
			AccessTime: time.Now(),
			ChangeTime: time.Now(),
			Name:       f,
			Mode:       0644,
			Size:       int64(len(fileData)),
		}
		tw.WriteHeader(&h)
		tw.Write(fileData)
		file.Close()
	}
	tw.Flush()
	tw.Close()
}
