package linked_list

import (
	"testing"

	"muse/data_structures/tests"
)

func TestLinkedList(t *testing.T) {
	t.Run("LinkedList", tests.ListDerive(New[int]))
}
