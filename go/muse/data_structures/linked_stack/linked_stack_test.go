package linked_stack

import (
	"testing"

	"muse/data_structures/tests"
)

func TestLinkedStack(t *testing.T) {
	t.Run("LinkedStack", tests.StackDerive(New[int]))
}
