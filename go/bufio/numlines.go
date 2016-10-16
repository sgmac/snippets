package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	filename = flag.String("f", "", "File to read")
	numbers  = flag.Bool("n", false, "Show line numbers")
)

func main() {

	flag.Parse()
	if flag.NFlag() < 1 {
		fmt.Fprintf(os.Stderr, "usage:%s -f <filename>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "-n  Show number lines.\n")
		os.Exit(1)
	}

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal("%s\n", err)
	}

	scanner := bufio.NewScanner(file)
	var lines int

	for scanner.Scan() {
		if *numbers {
			fmt.Fprintf(os.Stdout, "%d: %s\n", lines, scanner.Text())
			lines++
			continue
		}
		fmt.Fprintf(os.Stdout, "%s\n", scanner.Text())
	}
}
