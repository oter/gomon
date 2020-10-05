package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Test app returns code 1 on SIGTERM, 2 on SIGINT and 128 when 5-second timeout happens
func main() {
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	sigchan := make(chan os.Signal, len(signals))
	signal.Notify(sigchan, signals...)

	var exitCode int
	select {
	case sig := <-sigchan:
		switch sig {
		case syscall.SIGTERM:
			exitCode = 1
		case syscall.SIGINT:
			exitCode = 2
		}
	case <-time.After(5*time.Second):
		exitCode = 128
	}

	os.Exit(exitCode)
}
