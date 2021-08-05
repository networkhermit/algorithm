package TestRunner

import "testing"

func TestPick(t *testing.T) {
    for i := 0; i < 10; i++ {
        Pick(func() bool { return (i & 1) != 0 })
    }
}
