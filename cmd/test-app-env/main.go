package main

import "os"

// Test app returns code 0 on when env variable FOOBAR=bar is set
func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	v, ok := os.LookupEnv("FOOBAR")
	if !ok {
		exitCode = 1
		return
	}

	if v != "bar" {
		exitCode = 2
		return
	}
}
