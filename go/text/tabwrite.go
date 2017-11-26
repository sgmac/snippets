package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// writer, minwitdth, tabbing, padding
	w := tabwriter.NewWriter(os.Stdout, 20, 8, 0, ' ', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t33312499\t12\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()
}
