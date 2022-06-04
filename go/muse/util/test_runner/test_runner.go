package test_runner

import (
	"fmt"
	"os"
)

var itemIndex int

func Pick(fn func() bool) {
	if fn() {
		fmt.Printf("✓ Item [%d] PASSED\n", itemIndex)
	} else {
		fmt.Fprintf(os.Stderr, "✗ Item [%d] FAILED\n", itemIndex)
	}

	itemIndex++
}
