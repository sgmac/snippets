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

func processFile(file string, numlines int) error {
	var countLines int
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

loop:
	for _, r := range records {
		switch {
		case countLines < numlines:
			countLines += 1
		case countLines == numlines:
			break loop
		}
		fmt.Println(r)
	}
	return nil
}
func main() {
	file := flag.String("f", "", "Process file")
	numlines := flag.Int("n", -1, "Number of lines.")
	flag.Parse()

	if *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	processFile(*file, *numlines)
}
