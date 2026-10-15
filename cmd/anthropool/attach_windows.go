//go:build windows

package main

import (
	"fmt"
	"os"
)

func cmdAttach(args []string) int {
	fmt.Fprintln(os.Stderr, "anthropool: attach is not supported on Windows yet")
	return 1
}
