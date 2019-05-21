package TestRunner

import (
    "fmt"
    "os"
)

var TestRunner_TestIndex int

func ParseTest(ok bool) {
    if ok {
        fmt.Printf("✓ Test [%d] Passed\n", TestRunner_TestIndex)
    } else {
        fmt.Fprintf(os.Stderr, "✗ Test [%d] Failed\n", TestRunner_TestIndex)
    }

    TestRunner_TestIndex += 1
}
