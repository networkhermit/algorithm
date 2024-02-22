package test_runner

import "testing"

func TestPick(t *testing.T) {
	for i := range 10 {
		Pick(func() bool { return (i & 1) != 0 })
	}
}
