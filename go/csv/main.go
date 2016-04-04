package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func exitOnError(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func processFile(file string) error {

	b, err := ioutil.ReadFile(file)
	if err != nil {
		exitOnError(err)
	}

	reader := bytes.NewReader(b)
	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		exitOnError(err)
	}

	for _, r := range records {
		fmt.Println(r)
	}
	return nil
}
func main() {
	file := flag.String("f", "", "Process file")
	flag.Parse()

	if *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	processFile(*file)
}
