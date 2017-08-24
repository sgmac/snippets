package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// CustomTime holds a date
type CustomTime struct {
	Year  int
	Month int
	Day   int
}

func usage() string {

	out := `
usage: daysbet <date1> <date2>
       dateFormat  mm/dd/yyyy 
`
	return out
}

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%s\n", usage())
		os.Exit(2)
	}

	d1 := os.Args[1]
	d2 := os.Args[2]

	s1 := strings.Split(d1, "/")
	s2 := strings.Split(d2, "/")

	c1 := CustomTime{}
	c2 := CustomTime{}

	for i, v := range s1 {
		switch i {
		case 0:
			c1.Month, _ = strconv.Atoi(v)
		case 1:
			c1.Day, _ = strconv.Atoi(v)
		case 2:
			c1.Year, _ = strconv.Atoi(v)
		}
	}

	for i, v := range s2 {
		switch i {
		case 0:
			c2.Month, _ = strconv.Atoi(v)
		case 1:
			c2.Day, _ = strconv.Atoi(v)
		case 2:
			c2.Year, _ = strconv.Atoi(v)
		}
	}

	t1 := time.Date(c1.Year, time.Month(c1.Month), c1.Day, 0, 0, 0, 0, time.Local)
	t2 := time.Date(c2.Year, time.Month(c2.Month), c2.Day, 0, 0, 0, 0, time.Local)

	duration := t2.Sub(t1)
	fmt.Printf("There are : %.0f days\n", duration.Hours()/24)

}
