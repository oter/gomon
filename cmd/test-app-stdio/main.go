package main

import (
	"fmt"
	"os"
)

// Test app
func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "stderr":
		_, err := fmt.Fprintln(os.Stderr, "hello from stderr")
		if err != nil {
			panic(err)
		}
	case "stdout":
		_, err := fmt.Fprintln(os.Stdout, "hello from stdout")
		if err != nil {
			panic(err)
		}
	}
}
