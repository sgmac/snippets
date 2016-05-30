package greetings

import (
	"fmt"
	"os"
)

// Greetings receives one or more messages and sends them to stdout.
func Greetings(msg ...string) {
	fmt.Fprintf(os.Stdout, "%s\n", msg)
}
