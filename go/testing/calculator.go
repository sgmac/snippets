package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var ErrDivisionByZero = errors.New("Division by zero.")

// AddNumbers takes 2 numbers and returns the sum.
func AddNumbers(a, b int) int {
	return a + b
}

// AddNumbers takes 2 numbers and returns the difference.
func RestNumbers(a, b int) int {
	return a - b
}

// DivNumbers takes 2 numbers and performs integer division.
func DivNumbers(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func usage() {
	var help = `calculator: [OPTIONS] arg1 arg2
 -a:	Addition.
 -r:	Substraction.
 -d:	Division.
	`

	fmt.Fprintf(os.Stderr, "%s\n", help)
	os.Exit(1)
}
func main() {
	var a, b int
	flag.Usage = usage
	addition := flag.Bool("a", false, "")
	division := flag.Bool("d", false, "")
	substraction := flag.Bool("s", false, "")
	flag.Parse()

	if flag.NArg() < 2 {
		usage()
	}

	aStr := flag.Args()[0]
	bStr := flag.Args()[1]
	a, _ = strconv.Atoi(aStr)
	b, _ = strconv.Atoi(bStr)

	switch {
	case *addition:
		fmt.Printf("Result addition: %d\n", AddNumbers(a, b))
	case *division:
		result, err := DivNumbers(a, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Result division: %d\n", result)
	case *substraction:
		fmt.Printf("Result subsraction: %d\n", RestNumbers(a, b))
	}
}
