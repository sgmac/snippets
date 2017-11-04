package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6, 7, 8}
	s3 := []int{9, 10}

	var header, line string

	category := []string{"A", "B", "C"}

	w := tabwriter.NewWriter(os.Stdout, 40, 8, 0, ' ', 0)
	for i, e := range category {
		header = header + fmt.Sprintf("%s", e)
		if i < (len(category) - 1) {
			header = header + "\t"
		}
	}

	fmt.Fprintln(w, header)
	for i, _ := range s2 {
		if i < len(s2) {
			// Print values from slice1 when has values inside its own array, otherwise
			// just tab. Same for slice3.
			if i < len(s1) {
				line = line + strconv.Itoa(s1[i]) + "\t"
			} else {
				line = line + "\t"
			}

			line = line + strconv.Itoa(s2[i]) + "\t"

			if i < len(s3) {
				line = line + strconv.Itoa(s3[i])
			} else {
				line = line + "\t"
			}

		}
		fmt.Fprintln(w, line)
		line = ""
	}

	w.Flush()

}
