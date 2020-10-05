package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Test app receives SIGUSR1 signals.
// Exits with code 0 when receives SIGUSR1, otherwise with code 1 after 5-second timeout.
func main() {
	signals := []os.Signal{syscall.SIGUSR1}
	sigchan := make(chan os.Signal, len(signals))
	signal.Notify(sigchan, signals...)

	for {
		select {
		case sig := <-sigchan:
			switch sig {
			case syscall.SIGUSR1:
				return
			}
		case <-time.After(5*time.Second):
			os.Exit(1)
		}
	}
}
