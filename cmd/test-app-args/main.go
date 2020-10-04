package main

import "os"

// Test app expects either no cmd arguments or arguments from the next list: "-arg1", "-arg2", "-arg3".
// Returns the number of arguments passed if all the arguments are from the given list,
// otherwise returns 127.
func main() {
	validArguments := map[string]struct{}{
		"-arg1": {},
		"-arg2": {},
		"-arg3": {},
	}

	for _, arg := range os.Args[1:] {
		if _, ok := validArguments[arg]; !ok {
			os.Exit(127)
		}
	}

	os.Exit(len(os.Args[1:]))
}
