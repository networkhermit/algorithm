package array_stack

import (
	"testing"

	"muse/data_structures/tests"
)

func TestArrayStack(t *testing.T) {
	t.Run("ArrayStack(", tests.StackDerive(New[int]))
}
