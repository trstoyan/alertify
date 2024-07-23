package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Run the main function as a separate goroutine
	go main()

	// Perform any necessary setup or assertions here

	// Call os.Exit with the exit code returned by main
	os.Exit(m.Run())
}
