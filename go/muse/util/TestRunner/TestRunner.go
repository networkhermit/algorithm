package TestRunner

import (
    "fmt"
    "os"
)

var TestRunnerItemIndex int

func Pick(fn func() bool) {
    if fn() {
        fmt.Printf("✓ Item [%d] PASSED\n", TestRunnerItemIndex)
    } else {
        fmt.Fprintf(os.Stderr, "✗ Item [%d] FAILED\n", TestRunnerItemIndex)
    }

    TestRunnerItemIndex++
}
