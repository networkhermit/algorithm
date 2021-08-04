package TestRunner

import "testing"

func TestParseTest(t *testing.T) {
    for i := 0; i < 10; i++ {
        ParseTest((i & 1) != 0)
    }
}
