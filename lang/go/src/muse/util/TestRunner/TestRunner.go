package TestRunner

import (
    "fmt"
    "os"
)

var TestRunnerItemIndex int

func ParseTest(ok bool) {
    if ok {
        fmt.Printf("✓ Item [%d] PASSED\n", TestRunnerItemIndex)
    } else {
        fmt.Fprintf(os.Stderr, "✗ Item [%d] FAILED\n", TestRunnerItemIndex)
    }

    TestRunnerItemIndex++
}
