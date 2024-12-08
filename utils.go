package main

import (
	"fmt"
	"os"
)

// getPort returns the provided port Flag as a string
func getPort(port int) string {
	return fmt.Sprintf(":%d", port)
}

// handleErr writes the provided error to stderr
func handleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error occured: %v\n", err)
		os.Exit(1)
	}
}
