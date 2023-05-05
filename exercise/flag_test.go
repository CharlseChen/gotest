package exercise

import (
	"flag"
	"testing"
)

func FlagTesting(t *testing.T) {
	var firstVar int
	flag.IntVar(&firstVar, "n", 1234, "help message for flag n")
	flag.Parse()
	t.Skip("n:", firstVar)
}
